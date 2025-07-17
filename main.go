package main

import (
	"log"

	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
	"github.com/rohanthewiz/serr"
)

func main() {
	// Initialize the server with options
	s := rweb.NewServer(rweb.ServerOptions{
		Address: ":8080",
		Verbose: true,
	})

	// Add request info middleware for stats
	s.Use(rweb.RequestInfo)
	s.ElementDebugRoutes()

	// Define routes
	s.Get("/", rootHandler)
	s.Get("/about", aboutHandler)
	s.Get("/contact", contactHandler)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(s.Run())
}

// rootHandler serves the home page
func rootHandler(c rweb.Context) error {
	err := c.WriteHTML(generateHomePage())
	if err != nil {
		return serr.Wrap(err, "failed to write home page HTML")
	}
	return nil
}

// aboutHandler serves the about page
func aboutHandler(c rweb.Context) error {
	err := c.WriteHTML(generateAboutPage())
	if err != nil {
		return serr.Wrap(err, "failed to write about page HTML")
	}
	return nil
}

// contactHandler serves the contact page
func contactHandler(c rweb.Context) error {
	err := c.WriteHTML(generateContactPage())
	if err != nil {
		return serr.Wrap(err, "failed to write contact page HTML")
	}
	return nil
}

// generateHomePage creates the HTML for the home page
func generateHomePage() string {
	b := element.NewBuilder()

	b.Html().R(
		b.Head().R(
			b.Title().T("Welcome - My Basic Website"),
			b.Style().T(getStyles()),
		),
		b.Body().R(
			// Navigation component
			navigationComponent{}.Render(b),

			// Main content
			b.DivClass("container").R(
				b.H1().T("Welcome to My Website"),
				b.DivClass("content").R(
					b.P().T("This is a basic website built with Go using the rweb server and element HTML generation library."),
					b.P().T("Explore our pages to learn more about what we offer."),

					// Feature cards
					b.DivClass("features").R(
						b.DivClass("card").R(
							b.H3().T("Fast & Efficient"),
							b.P().T("Built with Go for optimal performance and resource efficiency."),
						),
						b.DivClass("card").R(
							b.H3().T("Clean Code"),
							b.P().T("Using element library for type-safe HTML generation with a code-first approach."),
						),
						b.DivClass("card").R(
							b.H3().T("Modern Design"),
							b.P().T("Simple, responsive design that works on all devices."),
						),
					),
				),
			),

			// Footer
			footerComponent{}.Render(b),
		),
	)

	return b.String()
}

// generateAboutPage creates the HTML for the about page
func generateAboutPage() string {
	b := element.NewBuilder()

	b.Html().R(
		b.Head().R(
			b.Title().T("About - My Basic Website"),
			b.Style().T(getStyles()),
		),
		b.Body().R(
			navigationComponent{}.Render(b),

			b.DivClass("container").R(
				b.H1().T("About Us"),
				b.DivClass("content").R(
					b.P().T("We are passionate about creating efficient web applications using modern Go technologies."),
					b.H2().T("Our Mission"),
					b.P().T("To demonstrate the power and simplicity of building web applications with Go, showcasing clean code practices and efficient server-side rendering."),
					b.H2().T("Technologies We Use"),
					b.Ul().R(
						b.Li().T("Go Programming Language"),
						b.Li().T("rweb - Lightweight web server framework"),
						b.Li().T("element - Type-safe HTML generation library"),
						b.Li().T("serr - Error handling library"),
					),
				),
			),

			footerComponent{}.Render(b),
		),
	)

	return b.String()
}

// generateContactPage creates the HTML for the contact page
func generateContactPage() string {
	b := element.NewBuilder()

	b.Html().R(
		b.Head().R(
			b.Title().T("Contact - My Basic Website"),
			b.Style().T(getStyles()),
		),
		b.Body().R(
			navigationComponent{}.Render(b),

			b.DivClass("container").R(
				b.H1().T("Contact Us"),
				b.DivClass("content").R(
					b.P().T("Get in touch with us through the following channels:"),
					b.DivClass("contact-info").R(
						b.P().R(
							b.Strong().T("Email: "),
							b.T("info@example.com"),
						),
						b.P().R(
							b.Strong().T("Phone: "),
							b.T("+1 (555) 123-4567"),
						),
						b.P().R(
							b.Strong().T("Address: "),
							b.T("123 Main Street, City, State 12345"),
						),
					),

					// Simple contact form
					b.H2().T("Send us a message"),
					b.Form("method", "post", "action", "#").R(
						b.DivClass("form-group").R(
							b.Label("for", "name").T("Name:"),
							b.Input("type", "text", "id", "name", "name", "name", "required", "required"),
						),
						b.DivClass("form-group").R(
							b.Label("for", "email").T("Email:"),
							b.Input("type", "email", "id", "email", "name", "email", "required", "required"),
						),
						b.DivClass("form-group").R(
							b.Label("for", "message").T("Message:"),
							b.TextArea("id", "message", "name", "message", "rows", "5", "required", "required").R(),
						),
						b.Button("type", "submit").T("Send Message"),
					),
				),
			),

			footerComponent{}.Render(b),
		),
	)

	return b.String()
}

// navigationComponent is a reusable navigation bar
type navigationComponent struct{}

func (n navigationComponent) Render(b *element.Builder) (x any) {
	b.Nav().R(
		b.DivClass("nav-container").R(
			b.DivClass("logo").T("My Website"),
			b.Ul().R(
				b.Li().R(b.A("href", "/").T("Home")),
				b.Li().R(b.A("href", "/about").T("About")),
				b.Li().R(b.A("href", "/contact").T("Contact")),
			),
		),
	)
	return
}

// footerComponent is a reusable footer
type footerComponent struct{}

func (f footerComponent) Render(b *element.Builder) (x any) {
	b.Footer().R(
		b.DivClass("footer-content").R(
			b.P().T("© 2024 My Basic Website. All rights reserved."),
			b.P().R(
				b.A("href", "#").T("Privacy Policy"),
				b.T(" | "),
				b.A("href", "#").T("Terms of Service"),
			),
		),
	)
	return
}

// getStyles returns the CSS styles for the website
func getStyles() string {
	return `
		/* Reset and base styles */
		* {
			margin: 0;
			padding: 0;
			box-sizing: border-box;
		}

		body {
			font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
			line-height: 1.6;
			color: #333;
			background-color: #f5f5f5;
		}

		/* Navigation styles */
		nav {
			background-color: #2c3e50;
			color: white;
			padding: 1rem 0;
			box-shadow: 0 2px 4px rgba(0,0,0,0.1);
		}

		.nav-container {
			max-width: 1200px;
			margin: 0 auto;
			padding: 0 2rem;
			display: flex;
			justify-content: space-between;
			align-items: center;
		}

		.logo {
			font-size: 1.5rem;
			font-weight: bold;
		}

		nav ul {
			list-style: none;
			display: flex;
			gap: 2rem;
		}

		nav a {
			color: white;
			text-decoration: none;
			transition: opacity 0.3s;
		}

		nav a:hover {
			opacity: 0.8;
		}

		/* Container and content */
		.container {
			max-width: 1200px;
			margin: 0 auto;
			padding: 2rem;
			min-height: calc(100vh - 200px);
		}

		h1 {
			color: #2c3e50;
			margin-bottom: 2rem;
			font-size: 2.5rem;
		}

		h2 {
			color: #34495e;
			margin: 2rem 0 1rem;
			font-size: 2rem;
		}

		h3 {
			color: #34495e;
			margin-bottom: 0.5rem;
		}

		.content {
			background: white;
			padding: 2rem;
			border-radius: 8px;
			box-shadow: 0 2px 4px rgba(0,0,0,0.1);
		}

		p {
			margin-bottom: 1rem;
		}

		/* Feature cards */
		.features {
			display: grid;
			grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
			gap: 2rem;
			margin-top: 2rem;
		}

		.card {
			background: #f8f9fa;
			padding: 1.5rem;
			border-radius: 8px;
			border: 1px solid #e9ecef;
			transition: transform 0.3s, box-shadow 0.3s;
		}

		.card:hover {
			transform: translateY(-4px);
			box-shadow: 0 4px 8px rgba(0,0,0,0.1);
		}

		/* Contact form */
		.contact-info {
			background: #f8f9fa;
			padding: 1.5rem;
			border-radius: 8px;
			margin-bottom: 2rem;
		}

		form {
			margin-top: 1rem;
		}

		.form-group {
			margin-bottom: 1.5rem;
		}

		label {
			display: block;
			margin-bottom: 0.5rem;
			font-weight: bold;
			color: #555;
		}

		input, textarea {
			width: 100%;
			padding: 0.75rem;
			border: 1px solid #ddd;
			border-radius: 4px;
			font-size: 1rem;
			font-family: inherit;
		}

		input:focus, textarea:focus {
			outline: none;
			border-color: #3498db;
			box-shadow: 0 0 0 2px rgba(52, 152, 219, 0.2);
		}

		button {
			background-color: #3498db;
			color: white;
			padding: 0.75rem 2rem;
			border: none;
			border-radius: 4px;
			font-size: 1rem;
			cursor: pointer;
			transition: background-color 0.3s;
		}

		button:hover {
			background-color: #2980b9;
		}

		/* Lists */
		ul {
			margin: 1rem 0 1rem 2rem;
		}

		/* Footer */
		footer {
			background-color: #2c3e50;
			color: white;
			padding: 2rem 0;
			text-align: center;
			margin-top: 3rem;
		}

		.footer-content {
			max-width: 1200px;
			margin: 0 auto;
			padding: 0 2rem;
		}

		footer a {
			color: #3498db;
			text-decoration: none;
		}

		footer a:hover {
			text-decoration: underline;
		}

		/* Responsive design */
		@media (max-width: 768px) {
			.nav-container {
				flex-direction: column;
				gap: 1rem;
			}

			nav ul {
				gap: 1rem;
			}

			h1 {
				font-size: 2rem;
			}

			.container {
				padding: 1rem;
			}

			.content {
				padding: 1rem;
			}
		}
	`
}
