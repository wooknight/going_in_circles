
BEGIN;

SET client_encoding = 'LATIN1';

CREATE TABLE basics(
    tconst char(9) PRIMARY KEY,
    titleType text,
    primaryTitle text,
    originalTitle text,
    isAdult integer,
    startYear integer,
    endYear integer,
    runtimeMinutes integer,
    genres text[]
) ;

COPY basics (tconst ,titleType ,primaryTitle,originalTitle ,isAdult ,startYear,endYear,runtimeMinutes,genres) FROM stdin;
tconst	titleType	primaryTitle	originalTitle	isAdult	startYear	endYear	runtimeMinutes	genres
tt0000001	short	Carmencita	Carmencita	0	1894	0	1	Documentary,Short
tt0000002	short	Le clown et ses chiens	Le clown et ses chiens	0	1892	0	5	Animation,Short
COMMIT;
