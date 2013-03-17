package opml

import (
    "encoding/xml"
    "io/ioutil"
    "testing"
)

func TestParse(t *testing.T) {
    body, err := ioutil.ReadFile("./greadersubscriptions.xml")
    if err != nil {
        t.Errorf("Error reading greadersubscriptions file %s", err)
    }
    subs := Opml1{}

    err = xml.Unmarshal(body, &subs)
    if err != nil {
        t.Errorf("Error unmarshalling greader subscriptions: %s", err)
    }

    if subs.Head.Title != "lamaofruin subscriptions in Google Reader" {
        t.Error("OPML title does not match")
    }

    if subs.Version != "1.0" {
        t.Error("OPML version mismatch")
    }

    greaderHead := OpmlHead{
        XMLName: xml.Name{Local: "head"},
        Title:   "lamaofruin subscriptions in Google Reader",
    }

    if subs.Head != greaderHead {
        t.Errorf("OPML GReader Head Mismatch - %+v - %+v", subs.Head, greaderHead)
    }

    greaderOutlines := []OpmlOutline{
        OpmlOutline{
            Text:    "(title unknown)",
            Title:   "(title unknown)",
            Type:    "rss",
            XmlUrl:  "http://pajoohyar.com/newsfeed",
            HtmlUrl: "http://pajoohyar.com/news",
        },
        OpmlOutline{
            Text:    "Ars Technica",
            Title:   "Ars Technica",
            Type:    "rss",
            XmlUrl:  "http://arstechnica.com/members/45928ace31705899/feeds/everything.xml",
            HtmlUrl: "http://arstechnica.com",
        },
        OpmlOutline{
            Text:    "Ars Technica",
            Title:   "Ars Technica",
            Type:    "rss",
            XmlUrl:  "http://feeds.arstechnica.com/arstechnica/etc?format=xml",
            HtmlUrl: "http://arstechnica.com/index.php",
        },
        OpmlOutline{
            Text:    "Ars Technica Features",
            Title:   "Ars Technica Features",
            Type:    "rss",
            XmlUrl:  "http://arstechnica.com/members/45928ace31705899/feeds/features/index.xml",
            HtmlUrl: "http://arstechnica.com",
        },
        OpmlOutline{
            Text:    "citeproc-js feed",
            Title:   "citeproc-js feed",
            Type:    "rss",
            XmlUrl:  "https://bitbucket.org/fbennett/citeproc-js/rss?token=c566adb106f26292e5e57c1e8e5446b6",
            HtmlUrl: "https://bitbucket.org/",
        },
        OpmlOutline{
            Text:    "Computers And Society updates on arXiv.org",
            Title:   "Computers And Society updates on arXiv.org",
            Type:    "rss",
            XmlUrl:  "http://export.arxiv.org/rss/cs.CY",
            HtmlUrl: "http://arxiv.org/",
        },
        OpmlOutline{
            Text:    "cs.AI updates on arXiv.org",
            Title:   "cs.AI updates on arXiv.org",
            Type:    "rss",
            XmlUrl:  "http://export.arxiv.org/rss/cs.AI",
            HtmlUrl: "http://arxiv.org/",
        },
        OpmlOutline{
            Text:    "cs.DC updates on arXiv.org",
            Title:   "cs.DC updates on arXiv.org",
            Type:    "rss",
            XmlUrl:  "http://export.arxiv.org/rss/cs.DC",
            HtmlUrl: "http://arxiv.org/",
        },
        OpmlOutline{
            Text:    "cs.DL updates on arXiv.org",
            Title:   "cs.DL updates on arXiv.org",
            Type:    "rss",
            XmlUrl:  "http://export.arxiv.org/rss/cs.DL",
            HtmlUrl: "http://arxiv.org/",
        },
        OpmlOutline{
            Text:    "cs.IR updates on arXiv.org",
            Title:   "cs.IR updates on arXiv.org",
            Type:    "rss",
            XmlUrl:  "http://export.arxiv.org/rss/cs.IR",
            HtmlUrl: "http://arxiv.org/",
        },
        OpmlOutline{
            Text:    "cs.MA updates on arXiv.org",
            Title:   "cs.MA updates on arXiv.org",
            Type:    "rss",
            XmlUrl:  "http://export.arxiv.org/rss/cs.MA",
            HtmlUrl: "http://arxiv.org/",
        },
        OpmlOutline{
            Text:    "cs.NI updates on arXiv.org",
            Title:   "cs.NI updates on arXiv.org",
            Type:    "rss",
            XmlUrl:  "http://export.arxiv.org/rss/cs.NI",
            HtmlUrl: "http://arxiv.org/",
        },

        OpmlOutline{
            Text:  "Listen Subscriptions",
            Title: "Listen Subscriptions",
            Outlines: []OpmlOutline{
                OpmlOutline{
                    Text:    "Digital Campus",
                    Title:   "Digital Campus",
                    Type:    "rss",
                    XmlUrl:  "http://feeds.feedburner.com/digitalcampus",
                    HtmlUrl: "http://digitalcampus.tv",
                },
                OpmlOutline{
                    Text:    "NodeUp",
                    Title:   "NodeUp",
                    Type:    "rss",
                    XmlUrl:  "http://feeds.feedburner.com/NodeUp",
                    HtmlUrl: "http://nodeup.com",
                },
                OpmlOutline{
                    Text:    "Science Talk",
                    Title:   "Science Talk",
                    Type:    "rss",
                    XmlUrl:  "http://rss.sciam.com/sciam/science-talk",
                    HtmlUrl: "http://www.scientificamerican.com/podcast/",
                },
            },
        },
    }

    CompareOutlineArrays(greaderOutlines, subs.Body.Outlines, t)
    /*
       for ind, val := range greaderOutlines {
           ol := greaderOutlines[ind]
           switch {
           case val.Text != ol.Text:
               t.Errorf("Outline.Text mismatch %d", ind)
           case val.Title != ol.Title:
               t.Errorf("Outline.Title mismatch %d", ind)
           case val.Type != ol.Type:
               t.Errorf("Outline.Type mismatch %d", ind)
           case val.XmlUrl != ol.XmlUrl:
               t.Errorf("Outline.XmlUrl mismatch %d", ind)
           }
       }
    */
}

func CompareOutlineArrays(o1, o2 []OpmlOutline, t *testing.T) {
    for ind, val := range o1 {
        ol2 := o2[ind]
        switch {
        case val.Text != ol2.Text:
            t.Errorf("Outline.Text mismatch %d", ind)
        case val.Title != ol2.Title:
            t.Errorf("Outline.Title mismatch %d", ind)
        case val.Type != ol2.Type:
            t.Errorf("Outline.Type mismatch %d", ind)
        case val.XmlUrl != ol2.XmlUrl:
            t.Errorf("Outline.XmlUrl mismatch %d", ind)
        case val.HtmlUrl != ol2.HtmlUrl:
            t.Errorf("Outline.HtmlUrl mismatch %d", ind)
        case true:
            t.Logf("Outlines matched %d - %s - %s", ind, val.XmlUrl, ol2.XmlUrl)
        }

        if len(val.Outlines) > 0 {
            CompareOutlineArrays(val.Outlines, ol2.Outlines, t)
        }
    }
}
