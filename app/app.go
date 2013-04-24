// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
       "html/template"
       "io/ioutil"
       "net/http"
       "regexp"
       "os"
       "fmt"
)

type Page struct {
     Title string
     Body  []byte
}

func (p *Page) save() error {
     filename := p.Title + ".txt"
     return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
     filename := title + ".txt"
     body, err := ioutil.ReadFile(filename)
     if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
     p, err := loadPage(title)
     if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
             return
             }
             renderTemplate(w, "view", p)
}



func editHandler(w http.ResponseWriter, r *http.Request, title string) {
     p, err := loadPage(title)
     if err != nil {
        p = &Page{Title: title}
    }
    renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
     body := r.FormValue("body")
     p := &Page{Title: title, Body: []byte(body)}
     err := p.save()
     if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
              return
              }
              http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
     err := templates.ExecuteTemplate(w, tmpl+".html", p)
     if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

const lenPath = len("/view/")

var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
     return func(w http.ResponseWriter, r *http.Request) {
            title := r.URL.Path[lenPath:]
              if !titleValidator.MatchString(title) {
                                http.NotFound(w, r)
                                        return
                                            }
                                                fn(w, r, title)
                                                }
}




func main() {
     http.HandleFunc("/", index)
     http.HandleFunc("/view/", makeHandler(viewHandler))
     http.HandleFunc("/edit/", makeHandler(editHandler))
     http.HandleFunc("/save/", makeHandler(saveHandler))
     http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}


func index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, 
        `<html><h1><code>Welcome to Go Lang on <a href='http://developer.CloudBees.com'>CloudBees</a></code></h1>
        <code>Made with a <a href='https://github.com/CloudBees-community/golang-clickstart'>ClickStart</a></code><p>
        <p>
            <code>Create a new <a href="/view/test">wiki page</a><code>
        </p>

        <!-- this is used for first time display of clickstart - a template-->
    <div id="clickstart_content" style="display:none">
    <p>
      Congratulations, your <a href="#CS_docUrl"><span>#CS_name</span></a> application is now running.<br />
      To modify it, <a href="https://grandcentral.cloudbees.com/user/ssh_keys">
      upload your public key (for git) here</a> if you haven't already.
      <br>Then clone your project:
    </p>
    <div class="CB_codeSample">
      git clone #CS_source #CS_appName<br/>
          cd #CS_appName<br/>
          ---- do your magic edits ----<br/>
          git commit -m "This is now even better"<br/>
          git push origin master
    </div>
    <p>That is it ! This will trigger your build/deploy pipline and publish your change</p>
    <p>We have set up all the moving parts for you - the management urls can be found on the following urls:</p>
    <ul>
      <li><strong>App console:</strong> <a href="#CS_appManageUrl">#CS_appManageUrl</a></li>
      <li><strong>Jenkins Build System:</strong> <a href="#CS_jenkinsUrl">#CS_jenkinsUrl</a></li>
      <li><strong>Source repositories:</strong> <a href="#CS_forgeUrl">#CS_forgeUrl</a></li>
    </ul>
  </div>
  <script type="text/javascript" src="https://ajax.googleapis.com/ajax/libs/jquery/1.4.0/jquery.min.js"></script>
  <script type="text/javascript" src="https://s3.amazonaws.com/cloudbees-downloads/clickstart/clickstart_intro.js"></script>
  <!-- end clickstart intro section -->     
    </html>`)
}