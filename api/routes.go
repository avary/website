package api

import (
	"net/http"

	"github.com/ibilalkayy/Bloop/website/handlers/account"
	"github.com/ibilalkayy/Bloop/website/handlers/authors"
	"github.com/ibilalkayy/Bloop/website/handlers/billing"
	"github.com/ibilalkayy/Bloop/website/handlers/books"
	"github.com/ibilalkayy/Bloop/website/handlers/callback"
	"github.com/ibilalkayy/Bloop/website/handlers/cart"
	"github.com/ibilalkayy/Bloop/website/handlers/checkout"
	"github.com/ibilalkayy/Bloop/website/handlers/home"
	"github.com/ibilalkayy/Bloop/website/handlers/login"
	"github.com/ibilalkayy/Bloop/website/handlers/logout"
	"github.com/ibilalkayy/Bloop/website/handlers/loops"
	"github.com/ibilalkayy/Bloop/website/handlers/signup"
	"github.com/ibilalkayy/Bloop/website/middleware"
)

func SetupRoutes() {
	http.HandleFunc("/", middleware.ErrorHandling(home.HomePage))
	http.HandleFunc("/loops", middleware.ErrorHandling(loops.LoopsPage))
	http.HandleFunc("/signup", middleware.ErrorHandling(signup.SignupPage))
	http.HandleFunc("/login", middleware.ErrorHandling(login.LoginPage))
	http.HandleFunc("/logout", middleware.ErrorHandling(logout.LogoutPage))
	http.HandleFunc("/callback", middleware.ErrorHandling(callback.CallbackPage))

	http.HandleFunc("/account", middleware.AuthenticatedOnly(middleware.ErrorHandling(account.AccountPage)))
	http.HandleFunc("/authors", middleware.ErrorHandling(authors.AuthorsPage))
	http.HandleFunc("/billing", middleware.ErrorHandling(billing.BillingPage))

	http.HandleFunc("/books", middleware.ErrorHandling(books.BooksPage))
	http.HandleFunc("/cart", middleware.ErrorHandling(cart.CartPage))
	http.HandleFunc("/checkout", middleware.ErrorHandling(checkout.CheckoutPage))
	http.HandleFunc("/authors-signup", middleware.AuthenticatedOnly(middleware.ErrorHandling(authors.AuthorsSignupPage)))

	fileServer := http.FileServer(http.Dir("./css/"))
	http.Handle("/css/", http.StripPrefix("/css", fileServer))

	// http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("../frontend/pages"))))
}
