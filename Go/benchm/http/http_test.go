package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func TestRegularDownload(t *testing.T) {
	url := "https://bstock.com/bohannon"
	statusCode := 200
	t.Logf("Given that we are testing HTTP connection")
	{
		t.Logf("Given the need to test downloading content")
		{
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("\t%s\t Should be able to GET - %v", failed, err)
			}
			t.Logf("\t%s\t Should be able to GET", succeed)
			defer resp.Body.Close()
			if resp.StatusCode != statusCode {
				t.Fatalf("\t%s\t Expected Response - %d - Got Response - %d - Error - %v", failed, statusCode, resp.StatusCode, err)
			}
			t.Logf("\t%s\t Expected Response - %d - Got Response - %d", succeed, statusCode, resp.StatusCode)
		}

	}
}

func BenchmarkRegularDownload(b *testing.B) {
	url := "https://bstock.com/bohannon"
	for i := 0; i < b.N; i++ {
		resp, err := http.Get(url)
		if err != nil {
			b.Fatalf("\t%s\t Should be able to GET - %v", failed, err)
		}
		b.Logf("\t%s\t Should be able to GET", succeed)
		defer resp.Body.Close()
	}
}

//http://feeds.marketwatch.com/marketwatch/topstories/
//http://seekingalpha.com/feed.xml

var feed = `<?xml-stylesheet type="text/xsl" media="screen" href="/~d/styles/rss2full.xsl"?><?xml-stylesheet type="text/css" media="screen" href="http://feeds.marketwatch.com/~d/styles/itemcontent.css"?><rss version="2.0">
<channel>
<title>MarketWatch.com - Top Stories</title>
<link>http://www.marketwatch.com/</link>
<description>MarketWatch, a leading publisher of business and financial news, offers users up-to-the minute news, investment tools, and subscription products.</description>
<language>en-us</language>
<copyright>Copyright 2021, MarketWatch, Inc.</copyright>
<ttl>60</ttl>
<image>
<title>MarketWatch.com - Top Stories</title>
<url>https://www.marketwatch.com/rss/marketwatch.gif</url>
<link>http://www.marketwatch.com/</link>
</image>
<atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="self" type="application/rss+xml" href="http://feeds.marketwatch.com/marketwatch/topstories" /><feedburner:info xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" uri="marketwatch/topstories" /><atom10:link xmlns:atom10="http://www.w3.org/2005/Atom" rel="hub" href="http://pubsubhubbub.appspot.com/" /><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="https://add.my.yahoo.com/rss?url=http%3A%2F%2Ffeeds.marketwatch.com%2Fmarketwatch%2Ftopstories" src="http://us.i1.yimg.com/us.yimg.com/i/us/my/addtomyyahoo4.gif">Subscribe with My Yahoo!</feedburner:feedFlare><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="http://www.newsgator.com/ngs/subscriber/subext.aspx?url=http%3A%2F%2Ffeeds.marketwatch.com%2Fmarketwatch%2Ftopstories" src="http://www.newsgator.com/images/ngsub1.gif">Subscribe with NewsGator</feedburner:feedFlare><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="http://feeds.my.aol.com/add.jsp?url=http%3A%2F%2Ffeeds.marketwatch.com%2Fmarketwatch%2Ftopstories" src="http://o.aolcdn.com/favorites.my.aol.com/webmaster/ffclient/webroot/locale/en-US/images/myAOLButtonSmall.gif">Subscribe with My AOL</feedburner:feedFlare><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="http://www.bloglines.com/sub/http://feeds.marketwatch.com/marketwatch/topstories" src="http://www.bloglines.com/images/sub_modern11.gif">Subscribe with Bloglines</feedburner:feedFlare><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="http://www.netvibes.com/subscribe.php?url=http%3A%2F%2Ffeeds.marketwatch.com%2Fmarketwatch%2Ftopstories" src="//www.netvibes.com/img/add2netvibes.gif">Subscribe with Netvibes</feedburner:feedFlare><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="http://fusion.google.com/add?feedurl=http%3A%2F%2Ffeeds.marketwatch.com%2Fmarketwatch%2Ftopstories" src="http://buttons.googlesyndication.com/fusion/add.gif">Subscribe with Google</feedburner:feedFlare><feedburner:feedFlare xmlns:feedburner="http://rssnamespace.org/feedburner/ext/1.0" href="http://www.pageflakes.com/subscribe.aspx?url=http%3A%2F%2Ffeeds.marketwatch.com%2Fmarketwatch%2Ftopstories" src="http://www.pageflakes.com/ImageFile.ashx?instanceId=Static_4&amp;fileName=ATP_blu_91x17.gif">Subscribe with Pageflakes</feedburner:feedFlare><item>
<title>: Inclusion of these U.S. money managers means nearly half of all asset funds managed globally are linked to climate-change pledge</title>
<link>http://www.marketwatch.com/news/story.asp?guid=%7B20C05575-04D4-B545-74FF-9518595F2274%7D&amp;siteid=rss&amp;rss=1</link>
<description>U.S. asset managers Franklin Templeton and MFS Investment Management are among the 41 new signatories, representing $6.8 trillion in assets, joining the Net Zero Asset Managers initiative.&lt;div class="feedflare"&gt;
&lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=CHf9BJ3MnGI:nogGcMR6oFs:yIl2AUoC8zA"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?d=yIl2AUoC8zA" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=CHf9BJ3MnGI:nogGcMR6oFs:F7zBnMyn0Lo"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?i=CHf9BJ3MnGI:nogGcMR6oFs:F7zBnMyn0Lo" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=CHf9BJ3MnGI:nogGcMR6oFs:V_sGLiPBpWU"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?i=CHf9BJ3MnGI:nogGcMR6oFs:V_sGLiPBpWU" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=CHf9BJ3MnGI:nogGcMR6oFs:qj6IDK7rITs"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?d=qj6IDK7rITs" border="0"&gt;&lt;/img&gt;&lt;/a&gt;
&lt;/div&gt;&lt;img src="http://feeds.feedburner.com/~r/marketwatch/topstories/~4/CHf9BJ3MnGI" height="1" width="1" alt=""/&gt;</description>
<pubDate>Tue, 06 Jul 2021 22:18:00 GMT</pubDate>
<guid isPermaLink="false">{20C05575-04D4-B545-74FF-9518595F2274}</guid>
</item>
<item>
<title>: AMC stock drops after company decides to not ask stockholders if it can issue more shares</title>
<link>http://www.marketwatch.com/news/story.asp?guid=%7B20C05575-04D4-B545-74FC-41487F84F0DC%7D&amp;siteid=rss&amp;rss=1</link>
<description>Shares of AMC Entertainment Holdings Inc. fell again Tuesday, erasing earlier gains, after the movie theater chain disclosed that it will no longer ask for shareholder approval to sell more shares.&lt;div class="feedflare"&gt;
&lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=vLC9KK8Kl6A:bijDTq077mc:yIl2AUoC8zA"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?d=yIl2AUoC8zA" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=vLC9KK8Kl6A:bijDTq077mc:F7zBnMyn0Lo"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?i=vLC9KK8Kl6A:bijDTq077mc:F7zBnMyn0Lo" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=vLC9KK8Kl6A:bijDTq077mc:V_sGLiPBpWU"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?i=vLC9KK8Kl6A:bijDTq077mc:V_sGLiPBpWU" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=vLC9KK8Kl6A:bijDTq077mc:qj6IDK7rITs"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?d=qj6IDK7rITs" border="0"&gt;&lt;/img&gt;&lt;/a&gt;
&lt;/div&gt;&lt;img src="http://feeds.feedburner.com/~r/marketwatch/topstories/~4/vLC9KK8Kl6A" height="1" width="1" alt=""/&gt;</description>
<pubDate>Tue, 06 Jul 2021 21:20:00 GMT</pubDate>
<guid isPermaLink="false">{20C05575-04D4-B545-74FC-41487F84F0DC}</guid>
</item>
<item>
<title>The Ratings Game: Virgin Galactic downgraded after stock more than doubled in six weeks</title>
<link>http://www.marketwatch.com/news/story.asp?guid=%7B20C05575-04D4-B545-74FC-C25C7539CE66%7D&amp;siteid=rss&amp;rss=1</link>
<description>Virgin Galactic Holdings Inc. was downgraded by UBS analyst Myles Walton, who cited concerns over valuation as the stock has more than doubled since he turned bullish about six weeks ago.&lt;div class="feedflare"&gt;
&lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=Cm8bL0mNLAU:64Xfwa38qNc:yIl2AUoC8zA"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?d=yIl2AUoC8zA" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=Cm8bL0mNLAU:64Xfwa38qNc:F7zBnMyn0Lo"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?i=Cm8bL0mNLAU:64Xfwa38qNc:F7zBnMyn0Lo" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=Cm8bL0mNLAU:64Xfwa38qNc:V_sGLiPBpWU"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?i=Cm8bL0mNLAU:64Xfwa38qNc:V_sGLiPBpWU" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=Cm8bL0mNLAU:64Xfwa38qNc:qj6IDK7rITs"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?d=qj6IDK7rITs" border="0"&gt;&lt;/img&gt;&lt;/a&gt;
&lt;/div&gt;&lt;img src="http://feeds.feedburner.com/~r/marketwatch/topstories/~4/Cm8bL0mNLAU" height="1" width="1" alt=""/&gt;</description>
<pubDate>Tue, 06 Jul 2021 21:13:00 GMT</pubDate>
<guid isPermaLink="false">{20C05575-04D4-B545-74FC-C25C7539CE66}</guid>
</item>
<item>
<title>Personal Finance Daily: Why no OPEC+ deal on crude oil production is bad news for U.S. drivers and surprise medical bills often strike after childbirth — but help is on the way</title>
<link>http://www.marketwatch.com/news/story.asp?guid=%7B20C05575-04D4-B545-7500-31446CC8138E%7D&amp;siteid=rss&amp;rss=1</link>
<description>&lt;div class="feedflare"&gt;
&lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=5d8FtCuxc1c:V89hDoNHmjI:yIl2AUoC8zA"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?d=yIl2AUoC8zA" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=5d8FtCuxc1c:V89hDoNHmjI:F7zBnMyn0Lo"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?i=5d8FtCuxc1c:V89hDoNHmjI:F7zBnMyn0Lo" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=5d8FtCuxc1c:V89hDoNHmjI:V_sGLiPBpWU"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?i=5d8FtCuxc1c:V89hDoNHmjI:V_sGLiPBpWU" border="0"&gt;&lt;/img&gt;&lt;/a&gt; &lt;a href="http://feeds.marketwatch.com/~ff/marketwatch/topstories?a=5d8FtCuxc1c:V89hDoNHmjI:qj6IDK7rITs"&gt;&lt;img src="http://feeds.feedburner.com/~ff/marketwatch/topstories?d=qj6IDK7rITs" border="0"&gt;&lt;/img&gt;&lt;/a&gt;
&lt;/div&gt;&lt;img src="http://feeds.feedburner.com/~r/marketwatch/topstories/~4/5d8FtCuxc1c" height="1" width="1" alt=""/&gt;</description>
<pubDate>Tue, 06 Jul 2021 21:03:00 GMT</pubDate>
<guid isPermaLink="false">{20C05575-04D4-B545-7500-31446CC8138E}</guid>
</item></channel>
</rss>
`

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("content-type", "application/xml")
		fmt.Fprintf(w, feed)
	}
	return httptest.NewServer(http.HandlerFunc(f))
}

func TestMockGet(t *testing.T) {
	statusCode := http.StatusOK
	server := mockServer()
	defer server.Close()
	t.Logf("Given the need to test downloading mock content")
	{
		t.Logf("\tTest 0:\tWhen checking status %q for status code %d.", server.URL, statusCode)
		resp, err := http.Get(server.URL)
		if err != nil {
			t.Fatalf("\t%s\tShould be able to get data from %s- Error %v", failed, server.URL, err)
		}
		t.Logf("\t%s\tShould be able to get data from %s", succeed, server.URL)
		defer resp.Body.Close()
		if resp.StatusCode != statusCode {
			t.Fatalf("\t%s\tShould receive a %d status code : Got %v", failed, statusCode, resp.StatusCode)
		}
		t.Logf("\t%s\tShould receive a %d status code : Got %v", succeed, statusCode, resp.StatusCode)
		var d Document
		if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
			t.Fatalf("\t%s\tShould be able to unmarshal the response : %v", failed, err)
		}
		t.Logf("\t%s\tShould be able to unmarshal the response", succeed)

	}
}


func BenchmarkMock(b *testing.B){
	
}