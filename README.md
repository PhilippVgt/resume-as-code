# resume-as-code
The project aims at simplifying the process of updating and maintaining your resume using common exchange formats YAML and JSON, keeping resume content and PDF formatting strictly separate.

This projects combines 3 components:
1. Well-defined YAML and JSON-based definition of resume content
2. HTML, CSS and Javascript-based template for resume formatting
3. HTML to PDF converter for PDF file generation

# Usage
To use resume-as-code on your machine, do the following:
1. Make sure you have Golang installed: https://go.dev/learn/
2. Clone the repository and run the build command `go build` inside the project directory.
   - This will create a binary file in the project directory:
      - `resume-as-code` on Linux and Mac OS
      - `resume-as-code.exe` on Windows
3. Run the binary with 2 arguments: `./resume-as-code <Resume Definition in YAML or JSON> <Template directory>`
   - e.g. `./resume-as-code /path/to/my-resume.yaml templates/modern`
   - Both resume definition and template directory can be referenced by relative or absolute paths.
   - If you run within the project directory, you can use the default template located at [templates/modern](templates/modern).
4. The resume PDF will be created next to the resume definition YAML or JSON.

# Schema
The resume schema provides structured representations for the following details:
- Personal Data (Name, Contact Details, Address, Web Profiles, etc.)
- Work
- Education
- Projects
- Publications
- Awards
- Certificates
- Skill Sets
- Language Skills
- Interests

Refer to `schema/resume.yaml` and `schema/resume.json` for all available attributes.

### Page Setup
For adjusting the PDF page use:
```yaml
page:
  size: A4          # Page format, A4 and Letter are supported, defaults to A4
  margins:          # Page margins in mm, global defaults are applied if omitted
    top: 10
    bottom: 15
    left: 5
    right: 10
```

### Template Theme
Depending on the template, the background and accent (or highlight) color can be adjusted to personalize the result:
```yaml
theme:
  background: "#243E66"     # Background color if applicable to the template
  accent: "#305285"         # Accent color (e.g. for headlines) if applicable to the theme
```
Values must be defined in HEX color code. Surround values with quotes in YAML, otherwise they will be parsed as comments and ignored. 

### Name
You can either define first and last name in separate fields, or combine them in one field if splitting first and last name is not common in your culture. Templates usually prefer `firstName` and `lastName` if present and otherwise fallback to the `name`. 
```yaml
personal:
  name: Peter Parker
  firstName: Peter
  lastName: Parker
  ...
```

### Photo
You can optionally add a photo to your resume by providing either an absolute file path on your system:
```yaml
personal:
  photo: /my/local/pictures/resume_photo.jpg
```
Alternatively you can supply a publicly accessible web URL:
```yaml
personal:
  photo: https://example.com/resume_photo.png
```

## Sample

A basic example could look like this (see [sample.yaml](sample/sample.yaml)):
```yaml
version: 1.0.0
page:
  size: Letter # or A4
personal:
  firstName: Peter
  lastName: Parker
  role: Biochemical Scientist
  summary: |
    Lorem ipsum dolor sit amet, consectetur adipiscing elit.
    Donec rutrum, risus quis tincidunt aliquam, mauris dui varius ipsum,
    quis tempus ligula leo gravida arcu. Vestibulum ante ipsum primis
    in faucibus orci luctus et ultrices posuere cubilia curae;
    Donec auctor sollicitudin metus, non placerat ex varius at.
  contact:
    phone: (555) 555-1234
    email: peter.parker@example.com
    website: https://example.com
  address:
    street: 20 Ingram Street
    city: Flushing, NY
    postalCode: 11375
    country: USA
  profiles:
    - platform: GitHub
      user: peter.parker
      website: https://example.com/peter.parker
work:
  - name: Octavius Industries
    address: New York City, USA
    role: Research Assistant
    website: https://www.example.com
    from: 2008-04-01
    until: 2012-08-31
    summary: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec rutrum, risus quis tincidunt aliquam, mauris dui varius ipsum, quis tempus ligula leo gravida arcu.
    highlights:
      - Lorem ipsum dolor sit amet, consectetur adipiscing elit.
      - Donec rutrum, risus quis tincidunt aliquam, mauris dui varius ipsum, quis tempus ligula leo gravida arcu.
education:
  - institution: Empire State University (ESU)
    address: New York City, USA
    website: https://www.example.com
    field: Biochemistry 
    degree: PhD
    from: 2005-10-15
    until: 2008-02-15
    grade: GPA 4.0
    summary: Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec rutrum, risus quis tincidunt aliquam, mauris dui varius ipsum, quis tempus ligula leo gravida arcu.
skillSets:
  - title: Physical Abilities
    skills:
      - skill: Swinging
        level: 0.9
      - skill: Jumping
        level: 0.7
languages:
  - language: English
    proficiency: Native
    level: 1.0
  - language: Spanish
    proficiency: Elementary Proficiency
    level: 0.2
```

For the above example, the resulting PDF will look something like this: [sample.pdf](sample/sample.pdf)


# Templates
Existing templates can be found in the [templates](templates) folder. Any common web technology can be used in templates as long as Google Chrome browser can handle it.
The default [modern](templates/modern) template is built with pure HTML and CSS and the javascript-based [Ionicons](https://ionic.io/ionicons) icon font.

### Contribute
Feel free to contribute your own template! Follow any existing template to get started.

**Things to note:**
- You can split stylesheets, scripts and html into individual files, but placeholder variables are only filled in the [index.html](templates/modern/index.html).
- Additional files must be placed inside a sub-folder `res` inside the template directory.
- Placeholders inside the template follow the variable names in the Resume structs. Check [resume.go](app/model/resume.go) if you are unsure of the supported data and the variable names. 