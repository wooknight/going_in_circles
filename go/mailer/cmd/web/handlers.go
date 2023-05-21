package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/wooknight/going_in_circles/go/mailer/data"
)

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}
	email := r.Form.Get("email")
	pass := r.Form.Get("password")
	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		app.Session.Put(r.Context(), "error", "invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	validPass, err := user.PasswordMatches(pass)
	if err != nil || !validPass {
		if err == nil {
			msg := Message{
				To:      email,
				Subject: "Failed login attempt",
				Data:    "invalid login attempt",
			}
			app.sendEmail(msg)
		}
		app.Session.Put(r.Context(), "error", "invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	app.Session.Put(r.Context(), "userID", user.ID)
	app.Session.Put(r.Context(), "user", user)
	app.Session.Put(r.Context(), "flash", "successful login")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {
	//clean up session
	_ = app.Session.Destroy(r.Context())
	_ = app.Session.RenewToken(r.Context())
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}
	//create a user
	u := data.User{
		Email:     r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Password:  r.Form.Get("password"),
		Active:    0,
		IsAdmin:   0,
	}
	_, err = u.Insert(u)
	if err != nil {
		app.Session.Put(r.Context(), "error", "unable to create user")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		app.ErrorLog.Println(err)
		return
	}

	//send an activation email
	url := fmt.Sprintf("http://localhost:%s/activate?email=%s", webPort, u.Email)
	signedURL := GenerateTokenFromString(url)
	app.InfoLog.Println(signedURL)
	msg := Message{
		To:       u.Email,
		Subject:  "activate your account",
		Template: "confirmation-email",
		Data:     template.HTML(signedURL),
	}
	app.sendEmail(msg)
	app.Session.Put(r.Context(), "flash", "confirmation email sent : Check your email")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	//validate url
	url := r.RequestURI
	testURL := fmt.Sprintf("http://localhost:%s%s", webPort, url)
	okay := VerifyToken(testURL)
	if !okay {
		app.Session.Put(r.Context(), "error", "invalid token.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//activate the account
	u, err := app.Models.User.GetByEmail(r.URL.Query().Get("email"))
	if err != nil {
		app.Session.Put(r.Context(), "error", "No user found.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u.Active = 1
	err = u.Update()
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to update user.")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	app.Session.Put(r.Context(), "flash", "Account activated, you can now login.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

	//generate a invoice

	//send an email with attachments

	//subscribe the user to an account
}

func (app *Config) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	if !app.Session.Exists(r.Context(), "userID") {
		app.Session.Put(r.Context(), "error", "You must be logged in see this page")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
		return
	}
	plans, err := app.Models.Plan.GetAll()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}
	dataMap := make(map[string]any)
	dataMap["plans"] = plans
	app.render(w, r, "plans.page.gohtml", &TemplateData{
		Data: dataMap,
	})
}

func (app *Config) SubscribeToPlan(w http.ResponseWriter, r *http.Request) {
	//get the ID of the plan
	id := r.URL.Query().Get("id")
	planId, _ := strconv.Atoi(id)

	//get the plan
	plan, err := app.Models.User.GetOne(planId)
	if err != nil {
		app.Session.Put(r.Context(), "error", "unable to find plan")
		http.Redirect(w, r, "/members/plans", http.StatusSeeOther)
		return
	}
	//get the user
	user, ok := app.Session.Get(r.Context(), "user").(data.User)
	if !ok {
		app.Session.Put(r.Context(), "error", "log in first")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	//generate an invoice
	app.Wait.Add(1)
	go func() {
		defer app.Wait.Done()
		invoice, err := app.getInvoice(user, plan)
		if err != nil {
			app.Mailer.ErrorChan <- err
		}
	}()
	//send an email with the invoice attached
	//generate a manual
	//send an email with the manual attached
	//redirect
}

func (app *Config) getInvoice(u data.User, plan *data.Plan) (string, error) {
	return plan.PlanAmountFormatted, nil
}
