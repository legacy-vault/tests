// pointer_test_c.go.

// Here we are dealing with Golang's Pointers Handling.
// The main Aim is to feel the Difference between various Ways of Object
// Initialization in Go Language.

package main

import (
	"fmt"
	"time"
)

type ClassX struct {
	Name             string
	Age              int
	Hobbies          []string
	Date             time.Time
	NeighbourObjects []*ClassX
	Size             uint64
	SubObject        []ClassXX
	LargeTexts       []string
}

type ClassXX struct {
	FA uint64
	FB string
	FC bool
}

func main() {

	test1()
	fmt.Println()
	test2()

	return
}

// Variant I.
func test1() {
	var obj_1 *ClassX
	obj_1 = new(ClassX)
	initApp_1(obj_1)
	//fmt.Println(obj_1)
}

// Variant II.
func test2() {
	var obj_2 *ClassX
	initApp_2(&obj_2)
	//fmt.Println(obj_2)
}

func initApp_1(x *ClassX) {
	tmp := NewClassXObject()
	*x = *tmp
}

func initApp_2(x **ClassX) {
	tmp := NewClassXObject()
	*x = tmp
}

func NewClassXObject() *ClassX {
	x := new(ClassX)
	x.init()
	return x
}

func (o *ClassX) init() {
	o.Name = "John"
	o.Age = 123
	o.Hobbies = []string{
		"Singing",
		"Running",
		"Music Listening",
		"Music Composing",
		"Car Fixing",
		"Speaking with Cats",
	}
	o.Date = time.Now()
	o.NeighbourObjects = make([]*ClassX, 20)
	o.Size = 12345
	o.SubObject = []ClassXX{
		ClassXX{FA: 1, FB: "Testing1", FC: true},
		ClassXX{FA: 2, FB: "Testing2", FC: true},
		ClassXX{FA: 3, FB: "Testing3", FC: true},
	}
	o.LargeTexts = []string{
		`
This page contains a list of all 302 languages for which official Wikipedias have been created under the auspices of the Wikimedia Foundation. This list includes 10 Wikipedias that were closed and moved to the Wikimedia Incubator for further development, so there are a current total of 292 active Wikipedias. Content in other languages is being developed at the Wikimedia Incubator; languages which meet certain criteria can get their own wikis.
The table entries are ordered by current article count. Each entry gives the language name in English (linked to the English Wikipedia article for the language), its "local name" (i.e. in the language itself, linked to the article in that language's wiki), the language code used in the wiki's URL address and in interwiki links to it (linked to the local Main Page), as well as statistics on articles, edits, administrators, users, active users, and images (most linked to an appropriate local special page).
To start a Wikipedia in a new language, please see our language proposal policy and the Incubator manual. Note: Just adding a link here does not create a new Wikipedia, nor does it serve to request that one be created.
If a wiki becomes active and is not listed here, please post a notice on this article's talk page, including a link to all the relevant Wikipedia pages, and help promote the effort by announcing it on the Wikipedia-L mailing list, and at Wikimedia News.
The tables here are regularly completely overwritten by editors (using automatically gathered data from the Special:Statistics page of each wiki), so edits made to individual entries won't last long, and are therefore usually unnecessary. If something is wrong with an entry other than simply having slightly out of date statistics, post about it on the talk page.
Lists of Wikipedias by various criteria:

    List of Wikipedias by article count, users, file count and depth (Source) (updated daily)
    List of Wikipedias by edits per article and depth (updated daily)
    List of Wikipedias by language group, family, and macrofamily (updated daily)
    List of Wikipedias by speakers per article (updated daily)
    List of Wikipedias by sample of articles and expanded sample of articles (both updated monthly)
    List of Wikipedias having zero local media files
    List of some Wikipedias by date created
    List of largest wikis (not just Wikimedia wikis)
    Wikipedia milestones (+en:Wikipedia:Milestone statistics) (tracking of major article-count milestones)
    Wikimedia News (announcements and tracking of milestones for all Wikimedia projects)`,

		`
Wikipedia (/ˌwɪkɪˈpiːdiə/ (About this sound listen), /ˌwɪkiˈpiːdiə/ (About this sound listen) WIK-ih-PEE-dee-ə) is a multilingual, web-based, free-content encyclopedia project supported by the Wikimedia Foundation and based on a model of openly editable content. The name "Wikipedia" is a portmanteau of the words wiki (a technology for creating collaborative websites, from the Hawaiian word wiki, meaning "quick") and encyclopedia. Wikipedia's articles provide links designed to guide the user to related pages with additional information.
Wikipedia is written collaboratively by largely anonymous volunteers who write without pay. Anyone with Internet access can write and make changes to Wikipedia articles, except in limited cases where editing is restricted to prevent disruption or vandalism. Users can contribute anonymously, under a pseudonym, or, if they choose to, with their real identity. The fundamental principles by which Wikipedia operates are the five pillars. The Wikipedia community has developed many policies and guidelines to improve the encyclopedia; however, it is not a formal requirement to be familiar with them before contributing.
Since its creation in 2001, Wikipedia has grown rapidly into one of the largest reference websites, attracting 374 million unique visitors monthly as of September 2015.[1] There are about 72,000 active contributors working on more than 48,000,000 articles in 302 languages. As of today, there are 5,740,361 articles in English. Every day, hundreds of thousands of visitors from around the world collectively make tens of thousands of edits and create thousands of new articles to augment the knowledge held by the Wikipedia encyclopedia. (See the statistics page for more information.) People of all ages, cultures and backgrounds can add or edit article prose, references, images and other media here. What is contributed is more important than the expertise or qualifications of the contributor. What will remain depends upon whether the content is free of copyright restrictions and contentious material about living people, and whether it fits within Wikipedia's policies, including being verifiable against a published reliable source, thereby excluding editors' opinions and beliefs and unreviewed research. Contributions cannot damage Wikipedia because the software allows easy reversal of mistakes and many experienced editors are watching to help ensure that edits are cumulative improvements. Begin by simply clicking the Edit link at the top of any editable page!
Wikipedia is a live collaboration differing from paper-based reference sources in important ways. Unlike printed encyclopedias, Wikipedia is continually created and updated, with articles on historic events appearing within minutes, rather than months or years. Because everybody can help improve it, Wikipedia has become more comprehensive than any other encyclopedia. In addition to quantity, its contributors work on improving quality as well. Wikipedia is a work-in-progress, with articles in various stages of completion. As articles develop, they tend to become more comprehensive and balanced. Quality also improves over time as misinformation and other errors are removed or repaired. However, because anyone can click "edit" at any time and add stuff in, any article may contain undetected misinformation, errors, or vandalism. Awareness of this helps the reader to obtain valid information, avoid recently added misinformation (see Wikipedia:Researching with Wikipedia), and fix the article. 
		`,

		`
In addition to her current teaching position at Evergreen, Coontz has also taught at Kobe University in Japan and the University of Hawaii at Hilo. She won the Washington Governor's Writers Award in 1989 for her book The Social Origins of Private Life: A History of American Families. In 1995 she received the Dale Richmond Award from the American Academy of Pediatrics for her "outstanding contributions to the field of child development." She received the 2001-02 "Friend of the Family" award from the Illinois Council on Family Relations. In 2004, she received the first-ever "Visionary Leadership" Award from the Council on Contemporary Families.
Coontz studies the history of American families, marriage, and changes in gender roles. Her book The Way We Never Were argues against several common myths about families of the past, including the idea that the 1950s family was traditional or the notion that families used to rely solely on their own resources. Her book, Marriage, A History: How Love Conquered Marriage, traces the history of marriage from Anthony and Cleopatra (not a love story, she argues) to debates over same-sex marriage. Her newest book, about the wives and daughters of "The Greatest Generation," is A Strange Stirring: The Feminine Mystique and American Women at the Dawn of the 1960s.
Coontz has appeared on national television and radio programs, including Oprah, the Today Show, The Colbert Report and dozens of NPR shows. In addition, her work has been featured in newspapers and magazines, as well as in many academic and professional journals. She has testified about her research before the House Select Committee on Children, Youth and Families and addressed audiences across America, Europe, and Japan.
In the landmark United States Supreme Court case Obergefell v. Hodges, justices cited Coontz's book, Marriage, A History in their decision to grant marriage equality to same-sex couples.[3][citation needed] 
		`,
	}
}
