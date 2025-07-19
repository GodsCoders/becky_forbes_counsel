package home

import (
	_ "embed"

	"github.com/rohanthewiz/element"
	"github.com/rohanthewiz/rweb"
	"github.com/rohanthewiz/serr"
)

//go:embed home.css
var pageStyles string

// HomeHandler serves the home page
func HomeHandler(c rweb.Context) error {
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
			b.Link("href", "https://fonts.googleapis.com/css2?family=Charm:wght@400;700&display=swap", "rel", "stylesheet"),

			b.Style().T(pageStyles),
		),
		b.Body().R(
			// Navigation
			element.RenderComponents(b, navigationComponent{}),

			// Hero Section
			element.RenderComponents(b, heroComponent{}),

			// Becky's Counseling Philosophy Section
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

// philosophyComponent creates the Becky's Counseling Philosophy section
type philosophyComponent struct{}

func (p philosophyComponent) Render(b *element.Builder) (x any) {
	b.Section("class", "philosophy-section").R(
		b.DivClass("container").R(
			b.DivClass("philosophy-content").R(
				b.DivClass("philosophy-image").R(
					b.Img("src", "/img/becky_portrait.jpeg", "alt", "Becky Forbes"),
				),
				b.DivClass("philosophy-text").R(
					b.H2().T("Becky's Counseling Philosophy"),
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
				b.P().T("✉ counseling@beckyrhoten.com"),
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
