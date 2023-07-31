package templates

import "html/template"

// Attaching the base.html file to each template.
func MakeTemplate(path string) *template.Template {
	files := []string{path, "../website/html/base.html"}
	return template.Must(template.ParseFiles(files...))
}

// Header Templates
var (
	// Main templates
	HomeTmpl = MakeTemplate("../website/html/home.html")

	// User templates
	// SignupTmpl  = MakeTemplate("../website/html/signup.html")
	// LoginTmpl   = MakeTemplate("../website/html/login.html")
	// LogoutTmpl  = MakeTemplate("../website/html/logout.html")
	AccountTmpl = MakeTemplate("../website/html/account.html")

	// Service templates
	AuthorsSignupTmpl = MakeTemplate("../website/html/authors_signup.html")
	AuthorsTmpl       = MakeTemplate("../website/html/authors.html")
	BooksTmpl         = MakeTemplate("../website/html/books.html")
	LoopsTmpl         = MakeTemplate("../website/html/loops.html")

	// Payment templates
	CartTmpl     = MakeTemplate("../website/html/cart.html")
	CheckoutTmpl = MakeTemplate("../website/html/checkout.html")
	BillingTmpl  = MakeTemplate("../website/html/billing.html")
)
