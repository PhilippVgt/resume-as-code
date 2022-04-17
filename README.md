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
   - If you run within the project directory, you can use the default template located at `templates/modern`.
4. The resume PDF will be created next to the resume definition YAML or JSON

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
