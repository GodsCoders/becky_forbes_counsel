package main

import (
	"log"

	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
	"github.com/rohanthewiz/serr"
)

func main() {
	/*
		TODO
		- move css to own file
		- templatize menu
		- pick out some decent Google fonts
		- create hierarchy of pages and components
	*/

	// Initialize the server with options
	s := rweb.NewServer(rweb.ServerOptions{
		Address: ":8080",
		Verbose: true,
	})

	// Add request info middleware for stats
	s.Use(rweb.RequestInfo)
	s.ElementDebugRoutes()

	// Define routes
	s.Get("/", homeHandler)

	log.Println("Starting server on http://localhost:8080")
	log.Fatal(s.Run())
}

// homeHandler serves the home page
func homeHandler(c rweb.Context) error {
	err := c.WriteHTML(generateHomePage())
	if err != nil {
		return serr.Wrap(err, "failed to write home page HTML")
	}
	return nil
}

// generateHomePage creates the HTML for the home page
func generateHomePage() string {
	b := element.NewBuilder()

	b.Html().R(
		b.Head().R(
			b.Meta("charset", "UTF-8"),
			b.Title().T("Becky Forbes Counseling - Find Strength, Healing, and Hope"),
			b.Meta("name", "viewport", "content", "width=device-width, initial-scale=1.0"),
			b.Link("rel", "preconnect", "href", "https://fonts.googleapis.com"),
			b.Link("rel", "preconnect", "href", "https://fonts.gstatic.com", "crossorigin", ""),
			b.Link("href", "https://fonts.googleapis.com/css2?family=Open+Sans:wght@300;400;600;700&display=swap", "rel", "stylesheet"),
			b.Style().T(getStyles()),
		),
		b.Body().R(
			// Navigation
			element.RenderComponents(b, navigationComponent{}),

			// Hero Section
			element.RenderComponents(b, heroComponent{}),

			// Sarah's Counseling Philosophy Section
			element.RenderComponents(b, philosophyComponent{}),

			// Benefits of Therapy Section
			element.RenderComponents(b, benefitsComponent{}),

			// Call to Action Section
			element.RenderComponents(b, ctaComponent{}),

			// Footer
			element.RenderComponents(b, footerComponent{}),
		),
	)

	return b.String()
}

// navigationComponent creates the navigation bar
type navigationComponent struct{}

func (n navigationComponent) Render(b *element.Builder) (x any) {
	b.Nav().R(
		b.DivClass("nav-container").R(
			// Logo
			b.DivClass("logo").R(
				b.Img("src", "/logo.png", "alt", "Becky Forbes Counseling"),
			),
			// Navigation menu
			b.DivClass("nav-menu").R(
				b.A("href", "/", "class", "active").T("Home"),
				b.A("href", "/about").T("About"),
				b.DivClass("dropdown").R(
					b.A("href", "/services", "class", "dropdown-toggle").R(
						b.T("Services "),
						b.Span().T("▾"),
					),
					b.DivClass("dropdown-content").R(
						b.A("href", "/individual-counseling").T("Individual Counseling"),
						b.A("href", "/couples-counseling").T("Couples Counseling"),
						b.A("href", "/emdr-therapy").T("EMDR Therapy"),
						b.A("href", "/career-counseling").T("Career Counseling"),
						b.A("href", "/anger-management-individual").T("Anger Management Individual Counseling"),
						b.A("href", "/anger-management").T("Anger Management"),
					),
				),
				b.A("href", "/fees").T("Fees"),
				b.A("href", "/faqs").T("FAQS"),
				b.A("href", "/contact").T("Contact"),
				b.A("href", "/resources").T("Resources"),
			),
		),
	)
	return
}

// heroComponent creates the hero section with background image
type heroComponent struct{}

func (h heroComponent) Render(b *element.Builder) (x any) {
	b.DivClass("hero-section").R(
		b.DivClass("hero-content").R(
			b.H1().T("Find Strength, Healing, and Hope"),
			b.P().T("Becky Forbes Counseling"),
			b.A("href", "/contact", "class", "btn btn-primary").T("Get In Touch"),
		),
	)
	return
}

// philosophyComponent creates the Sarah's Counseling Philosophy section
type philosophyComponent struct{}

func (p philosophyComponent) Render(b *element.Builder) (x any) {
	b.Section("class", "philosophy-section").R(
		b.DivClass("container").R(
			b.DivClass("philosophy-content").R(
				b.DivClass("philosophy-image").R(
					b.Img("src", "/sarah-photo.jpg", "alt", "Sarah Rhoten"),
				),
				b.DivClass("philosophy-text").R(
					b.H2().T("Sarah's Counseling Philosophy"),
					b.P().T("As a therapist, I seek to empower my clients with the knowledge that they hold the key to their well-being and help them develop the skills they need to access this ability on a daily basis. My compassion, empathy, and love of people along with my own life journey is what compelled me to enter the field of counseling. My primary theoretical model is Cognitive Behavioral Therapy (CBT), combined with faith-based counseling. My counseling philosophy is from a holistic point of view addressing the mental, physical, and spiritual aspects of a person. I also believe therapy must be tailored to the individual as no two people are the same. I prefer to use a client-centered collaborative approach, focusing on individuals strengths to create positive change and personal growth."),
				),
			),
		),
	)
	return
}

// benefitsComponent creates the Benefits of Therapy section
type benefitsComponent struct{}

func (ben benefitsComponent) Render(b *element.Builder) (x any) {
	b.Section("class", "benefits-section").R(
		b.DivClass("container").R(
			b.H2().T("Benefits of Therapy"),
			b.DivClass("benefits-grid").R(
				// Mental Health column
				b.DivClass("benefit-column").R(
					b.H3().T("Mental Health"),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Self-awareness- Therapy can help you understand your thoughts, emotions, and behaviors, and how they may be affecting your life."),
					),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Self-esteem- Therapy can help you feel more confident and accept yourself, even if you have negative thoughts about yourself."),
					),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Stress management- Therapy can teach you effective ways to manage stress, which can improve your sleep, reduce blood pressure, and strengthen your immune system."),
					),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Trauma- Therapy can help you make sense of past trauma and overcome fears."),
					),
				),
				// Physical Health column
				b.DivClass("benefit-column").R(
					b.H3().T("Physical Health"),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Boost immune system- Therapy can improve your overall health. People with mental health issues such as anxiety and depression often have compromised immune systems. Seeking help for your mental health can help you live a healthier and longer life."),
					),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Pain relief- Therapy can help reduce pain."),
					),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Improved sleep- Therapy can help you sleep better and is proven to be helpful with those suffering from insomnia."),
					),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Reduced risk of heart disease- Therapy can help reduce mental exhaustion, which can reduce stress throughout your body and potentially reduce the risk of heart disease."),
					),
				),
				// Interpersonal Relationships/Life Changes column
				b.DivClass("benefit-column").R(
					b.H3().T("Interpersonal Relationships/Life Changes"),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Stronger relationships- Therapy can help improve relationships by providing individuals with the tools to communicate more effectively, set healthy boundaries, and work through conflicts."),
					),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Communication skills- Therapy can help you identify communication issues and develop skills to articulate your needs, listen actively, set healthy boundaries, and resolve conflicts."),
					),
					b.DivClass("benefit-item").R(
						b.Span("class", "checkmark").T("✓"),
						b.P().T("Life transitions- Therapy can help you adapt to life transitions, such as moving out, getting married, divorce/ending a relationship, job/career changes, death/loss."),
					),
				),
			),
		),
	)
	return
}

// ctaComponent creates the call to action section
type ctaComponent struct{}

func (c ctaComponent) Render(b *element.Builder) (x any) {
	b.Section("class", "cta-section").R(
		b.DivClass("container").R(
			b.H2().T("Start your journey today"),
			b.H3().T("with Becky Forbes Counseling"),
			b.A("href", "/contact", "class", "btn btn-secondary").T("Schedule Now"),
		),
	)
	return
}

// footerComponent creates the footer
type footerComponent struct{}

func (f footerComponent) Render(b *element.Builder) (x any) {
	b.Footer().R(
		b.DivClass("footer-content").R(
			b.DivClass("footer-section").R(
				b.P().T("✉ counseling@sarahrhoten.com"),
				b.P().T("☎ 817-123-4567"),
				b.P().T("☎ 817-890-1234"),
			),
			b.DivClass("footer-section").R(
				b.DivClass("social-links").R(
					b.A("href", "#", "class", "social-link").T("Instagram"),
					b.A("href", "#", "class", "social-link").T("LinkedIn"),
				),
			),
			b.DivClass("footer-section").R(
				b.Img("src", "/footer-logo.png", "alt", "Becky Forbes Counseling", "class", "footer-logo"),
			),
		),
		b.DivClass("footer-bottom").R(
			b.DivClass("container").R(
				b.P().T("Copyright © 2025 Becky Forbes Counseling. All rights reserved."),
				b.DivClass("footer-links").R(
					b.A("href", "/privacy").T("Privacy Policy"),
					b.T(" | "),
					b.A("href", "/terms").T("Terms & Conditions"),
				),
			),
		),
	)
	return
}

// getStyles returns the CSS styles for the website
func getStyles() string {
	// Design system colors matching the screenshot
	return `
		/* CSS Variables for theme colors */
		:root {
			/* Primary colors */
			--color-primary: #5a9e82;
			--color-primary-rgb: 90, 158, 130;
			
			/* Secondary colors */
			--color-secondary: #f0c419;
			--color-secondary-hover: #d4ab17;
			
			/* Text colors */
			--color-text-primary: #333;
			--color-text-secondary: #555;
			--color-text-light: #fff;
			
			/* Background colors */
			--color-bg-primary: #fff;
			--color-bg-secondary: #f8f9fa;
			--color-bg-tertiary: #f5f5f5;
			
			/* Shadow colors */
			--shadow-light: rgba(0, 0, 0, 0.1);
			--shadow-medium: rgba(0, 0, 0, 0.2);
			--shadow-dark: rgba(0, 0, 0, 0.4);
			
			/* Overlay opacity */
			--overlay-opacity: 0.9;
		}

		/* Reset and base styles */
		* {
			margin: 0;
			padding: 0;
			box-sizing: border-box;
		}

		body {
			font-family: 'Open Sans', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
			line-height: 1.6;
			color: var(--color-text-primary);
			background-color: var(--color-bg-primary);
		}

		/* Navigation styles */
		nav {
			background-color: var(--color-bg-primary);
			padding: 1rem 0;
			box-shadow: 0 2px 4px var(--shadow-light);
			position: sticky;
			top: 0;
			z-index: 100;
		}

		.nav-container {
			max-width: 1200px;
			margin: 0 auto;
			padding: 0 2rem;
			display: flex;
			justify-content: space-between;
			align-items: center;
		}

		.logo img {
			height: 60px;
			width: auto;
		}

		.nav-menu {
			display: flex;
			gap: 2rem;
			align-items: center;
		}

		.nav-menu > a,
		.nav-menu .dropdown > a {
			color: var(--color-text-primary);
			text-decoration: none;
			font-weight: 500;
			transition: all 0.3s ease;
			position: relative;
			padding-bottom: 0.5rem;
		}

		/* Cool underline effect */
		.nav-menu > a::after,
		.nav-menu .dropdown > a::after {
			content: '';
			position: absolute;
			bottom: 0;
			left: 50%;
			width: 0;
			height: 3px;
			background: var(--color-secondary);
			transition: all 0.3s ease;
			transform: translateX(-50%);
			border-radius: 2px;
			box-shadow: 0 0 0 rgba(240, 196, 25, 0);
		}

		.nav-menu > a:hover::after,
		.nav-menu .dropdown:hover > a::after {
			width: 100%;
			box-shadow: 0 2px 8px rgba(240, 196, 25, 0.4);
		}

		/* Active menu item */
		.nav-menu > a.active::after,
		.nav-menu .dropdown > a.active::after {
			width: 100%;
			background: var(--color-secondary);
			box-shadow: 0 2px 12px rgba(240, 196, 25, 0.6);
			animation: glow 2s ease-in-out infinite;
		}

		@keyframes glow {
			0%, 100% {
				box-shadow: 0 2px 12px rgba(240, 196, 25, 0.6);
			}
			50% {
				box-shadow: 0 2px 20px rgba(240, 196, 25, 0.8);
			}
		}

		.nav-menu > a:hover,
		.nav-menu .dropdown > a:hover {
			color: var(--color-primary);
			transform: translateY(-2px);
		}

		/* Dropdown menu styles */
		.dropdown {
			position: relative;
		}

		.dropdown-content {
			display: none;
			position: absolute;
			background-color: var(--color-bg-primary);
			min-width: 280px;
			box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
			z-index: 1000;
			border-radius: 8px;
			top: calc(100% + 0.5rem);
			left: 0;
			padding: 0.5rem;
			margin-top: -0.5rem;
			border: 1px solid rgba(90, 158, 130, 0.1);
		}

		.dropdown:hover .dropdown-content {
			display: block;
		}

		/* Keep dropdown open when hovering over submenu items */
		.dropdown-content:hover {
			display: block;
		}

		/* Add invisible bridge to prevent gap */
		.dropdown::after {
			content: '';
			position: absolute;
			top: 100%;
			left: 0;
			right: 0;
			height: 0.5rem;
			display: none;
		}

		.dropdown:hover::after {
			display: block;
		}

		.dropdown-content a {
			display: block;
			padding: 0.75rem 1rem;
			transition: all 0.3s ease;
			border-radius: 6px;
			margin: 0.2rem 0;
			color: var(--color-text-primary);
			text-decoration: none;
			font-weight: 400;
			position: relative;
			overflow: hidden;
		}

		.dropdown-content a::before {
			content: '';
			position: absolute;
			left: 0;
			top: 50%;
			transform: translateY(-50%);
			width: 3px;
			height: 0;
			background-color: var(--color-secondary);
			transition: height 0.3s ease;
		}

		.dropdown-content a:hover {
			background-color: rgba(90, 158, 130, 0.1);
			padding-left: 1.5rem;
			color: var(--color-primary);
		}

		.dropdown-content a:hover::before {
			height: 70%;
		}

		/* Hero section styles */
		.hero-section {
			background: linear-gradient(var(--shadow-dark), var(--shadow-dark)), url('/hero-bg.jpg');
			background-size: cover;
			background-position: center;
			background-attachment: fixed;
			min-height: 600px;
			display: flex;
			align-items: center;
			justify-content: center;
			text-align: center;
			color: var(--color-text-light);
		}

		.hero-content h1 {
			font-size: 3rem;
			margin-bottom: 1rem;
			font-weight: 300;
			letter-spacing: 1px;
		}

		.hero-content p {
			font-size: 1.5rem;
			margin-bottom: 2rem;
			font-weight: 300;
		}

		/* Button styles */
		.btn {
			display: inline-block;
			padding: 1rem 2.5rem;
			text-decoration: none;
			border-radius: 4px;
			transition: all 0.3s;
			font-weight: 600;
			text-transform: uppercase;
			letter-spacing: 1px;
		}

		.btn-primary {
			background-color: var(--color-secondary);
			color: var(--color-text-primary);
		}

		.btn-primary:hover {
			background-color: var(--color-secondary-hover);
			transform: translateY(-2px);
			box-shadow: 0 4px 8px var(--shadow-medium);
		}

		.btn-secondary {
			background-color: var(--color-secondary);
			color: var(--color-text-primary);
			padding: 1.2rem 3rem;
			font-size: 1.1rem;
		}

		.btn-secondary:hover {
			background-color: var(--color-secondary-hover);
			transform: translateY(-2px);
			box-shadow: 0 4px 8px var(--shadow-medium);
		}

		/* Container */
		.container {
			max-width: 1200px;
			margin: 0 auto;
			padding: 0 2rem;
		}

		/* Philosophy section */
		.philosophy-section {
			padding: 5rem 0;
			background-color: var(--color-bg-secondary);
		}

		.philosophy-content {
			display: grid;
			grid-template-columns: 300px 1fr;
			gap: 3rem;
			align-items: start;
		}

		.philosophy-image img {
			width: 100%;
			border-radius: 8px;
			box-shadow: 0 4px 8px var(--shadow-light);
		}

		.philosophy-text h2 {
			color: var(--color-primary);
			font-size: 2.5rem;
			margin-bottom: 1.5rem;
			font-style: italic;
			font-weight: 400;
		}

		.philosophy-text p {
			font-size: 1.1rem;
			line-height: 1.8;
			color: var(--color-text-secondary);
		}

		/* Benefits section */
		.benefits-section {
			padding: 5rem 0;
			background-color: var(--color-bg-primary);
		}

		.benefits-section h2 {
			text-align: center;
			color: var(--color-primary);
			font-size: 2.5rem;
			margin-bottom: 3rem;
			font-style: italic;
			font-weight: 400;
		}

		.benefits-grid {
			display: grid;
			grid-template-columns: repeat(3, 1fr);
			gap: 3rem;
		}

		.benefit-column h3 {
			color: var(--color-primary);
			font-size: 1.5rem;
			margin-bottom: 1.5rem;
			padding-bottom: 1rem;
			border-bottom: 2px solid var(--color-primary);
		}

		.benefit-item {
			display: flex;
			align-items: start;
			margin-bottom: 1.5rem;
			gap: 1rem;
		}

		.checkmark {
			color: var(--color-primary);
			font-size: 1.2rem;
			flex-shrink: 0;
			margin-top: 0.2rem;
		}

		.benefit-item p {
			font-size: 0.95rem;
			line-height: 1.6;
			color: var(--color-text-secondary);
		}

		/* CTA section */
		.cta-section {
			padding: 5rem 0;
			background: linear-gradient(rgba(var(--color-primary-rgb), var(--overlay-opacity)), rgba(var(--color-primary-rgb), var(--overlay-opacity))), url('/cta-bg.jpg');
			background-size: cover;
			background-position: center;
			text-align: center;
			color: var(--color-text-light);
		}

		.cta-section h2 {
			font-size: 2.5rem;
			margin-bottom: 0.5rem;
			font-weight: 300;
			font-style: italic;
		}

		.cta-section h3 {
			font-size: 2rem;
			margin-bottom: 2rem;
			font-weight: 300;
			font-style: italic;
		}

		/* Footer */
		footer {
			background-color: var(--color-primary);
			color: var(--color-text-light);
			padding: 3rem 0 0;
		}

		.footer-content {
			max-width: 1200px;
			margin: 0 auto;
			padding: 0 2rem;
			display: grid;
			grid-template-columns: repeat(3, 1fr);
			gap: 3rem;
			align-items: center;
			text-align: center;
		}

		.footer-section p {
			margin-bottom: 0.5rem;
		}

		.social-links {
			display: flex;
			gap: 2rem;
			justify-content: center;
		}

		.social-link {
			color: var(--color-text-light);
			text-decoration: none;
			font-size: 1.2rem;
			transition: opacity 0.3s;
		}

		.social-link:hover {
			opacity: 0.8;
		}

		.footer-logo {
			height: 80px;
			width: auto;
		}

		.footer-bottom {
			background-color: var(--color-secondary);
			color: var(--color-text-primary);
			padding: 1rem 0;
			margin-top: 2rem;
		}

		.footer-bottom .container {
			display: flex;
			justify-content: space-between;
			align-items: center;
		}

		.footer-links a {
			color: var(--color-text-primary);
			text-decoration: none;
			transition: opacity 0.3s;
		}

		.footer-links a:hover {
			opacity: 0.7;
		}

		/* Responsive design */
		@media (max-width: 768px) {
			.nav-menu {
				flex-direction: column;
				gap: 1rem;
			}

			.hero-content h1 {
				font-size: 2rem;
			}

			.philosophy-content {
				grid-template-columns: 1fr;
			}

			.benefits-grid {
				grid-template-columns: 1fr;
			}

			.footer-content {
				grid-template-columns: 1fr;
				gap: 2rem;
			}

			.footer-bottom .container {
				flex-direction: column;
				gap: 1rem;
				text-align: center;
			}
		}
	`
}
