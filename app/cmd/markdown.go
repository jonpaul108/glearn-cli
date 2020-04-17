package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var markdownCmd = &cobra.Command{
	Use:     "markdown",
	Aliases: []string{"md"},
	Short:   "Copy curriculum markdown to clipboard",
	Long:    "Copy curriculum markdown to clipboard. Takes one argument, the type of content to copy to clipboard.\n\n" + argList,
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(incorrectNumArgs)
			os.Exit(1)
		}
		copyContent(args[0])
	},
}

type temp struct {
	Name     string
	Template string
}

func copyContent(command string) {
	t, ok := templates[command]

	if !ok {
		fmt.Println("Unknown arg " + command + ". Run 'learn md --help' for options.")
		return
	}

	if contains(noIdCommands, command) {
		clipboard.WriteAll(t.Template)
		fmt.Println(t.Name, "copied to clipboard!")
	} else {
		id := uuid.New().String()
		clipboard.WriteAll(fmt.Sprintf(strings.ReplaceAll(t.Template, `~~~`, "```"), id))
		fmt.Println(t.Name, "copied to clipboard!\nid:", id)
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var noIdCommands = []string{"ls", "lesson", "tpr", "testableproject", "cfy", "configyaml", "cry", "courseyaml"}

var templates = map[string]temp{
	"ls":              {"Lesson markdown", lessonTemplate},
	"lesson":          {"Lesson markdown", lessonTemplate},
	"mc":              {"Multiple choice markdown", multiplechoiceTemplate},
	"multiplechoice":  {"Multiple choice markdown", multiplechoiceTemplate},
	"cb":              {"Checkbox markdown", checkboxTemplate},
	"checkbox":        {"Checkbox markdown", checkboxTemplate},
	"sa":              {"Short answer markdown", shortanswerTemplate},
	"shortanswer":     {"Short answer markdown", shortanswerTemplate},
	"nb":              {"Number markdown", numberTemplate},
	"number":          {"Number markdown", numberTemplate},
	"pg":              {"Paragraph markdown", paragraphTemplate},
	"paragraph":       {"Paragraph markdown", paragraphTemplate},
	"js":              {"Javascript markdown", javascriptTemplate},
	"javascript":      {"Javascript markdown", javascriptTemplate},
	"ja":              {"Java markdown", javaTemplate},
	"java":            {"Java markdown", javaTemplate},
	"py":              {"Python markdown", pythonTemplate},
	"python":          {"Python markdown", pythonTemplate},
	"sq":              {"Sql markdown", sqlTemplate},
	"sql":             {"Sql markdown", sqlTemplate},
	"pr":              {"Project markdown", projectTemplate},
	"project":         {"Project markdown", projectTemplate},
	"tpr":             {"Testable Project markdown", testableProjectTemplate},
	"testableproject": {"Testable Project markdown", testableProjectTemplate},
	"cfy":             {"config.yaml syntax", configyamlTemplate},
	"configyaml":      {"config.yaml syntax", configyamlTemplate},
	"cry":             {"course.yaml syntax", courseyamlTemplate},
	"courseyaml":      {"course.yaml syntax", courseyamlTemplate},
}

const incorrectNumArgs = "Incorrect number of args. Takes one argument, the type of content to copy to clipboard.\n\n" + argList

const argList = `Args, full (abbreviation)--

Files:
  lesson (ls)
Questions:
  multiplechoice (mc)
  checkbox (cb)
  shortanswer (sa)
  number (nb)
  paragraph (pg)
  javascript (js)
  java (ja)
  python (py)
  sql (sq)
  project (pr)
  testableproject (tpr)
Configuration:
  configyaml (cfy)
  courseyaml (cry)`

const lessonTemplate = `# Title

## Learning Objectives

By the end of this lesson you will be able to:

* First Objective
* [at least one]
* [no more than four]

## Lesson Content

[Can be written content, videos, slides, images, gifs, etc. Think about including a rationale as the first few sentences/paragraph if you feel the lesson requires significant motivation or context. Examples of markdown formatting are at https://learn-2.galvanize.com/cohorts/667/blocks/13/content_files/walkthrough/03b-markdown-examples.md]

## Challenges

[It's recommended that each lesson has at least one challenge. Challenges make the content interactive and give instructors visibility into student learning. These challenge can be spread out in between content, or can be at the end of the lesson. Examples of all challenge types are in this unit -- https://learn-2.galvanize.com/cohorts/667/blocks/13/content_files/Multiple-Choice-Challenge.md]`

const multiplechoiceTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: multiple-choice
* id: %s
* title: [text, a short question title]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !options

* [Option 1]
* [Option 2]
* [Option 3, etc]

##### !end-options

##### !answer

* [Option 2 (the correct answer)]

##### !end-answer

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const checkboxTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: checkbox
* id: %s
* title: [text, a short question title]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !options

* [Option 1]
* [Option 2]
* [Option 3, etc]

##### !end-options

##### !answer

* [Option 2]
* [Option 3 (the correct answer set)]

##### !end-answer

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const shortanswerTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: short-answer
* id: %s
* title: [text, a short question title]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !placeholder

[text, placeholder text for the input field]

##### !end-placeholder

##### !answer

[text or regex, the answer (if regex wrap in /)]

##### !end-answer

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const numberTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: number
* id: %s
* title: [text, a short question title]
* decimal: [optional number, decimal points to user for answer evaluation]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !placeholder

[text, placeholder text for input field]

##### !end-placeholder

##### !answer

[number, the correct answer]

##### !end-answer

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const paragraphTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: paragraph
* id: %s
* title: [text, a short question title]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !placeholder

[text, placeholder text for input field]

##### !end-placeholder

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const javascriptTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: code-snippet
* language: javascript
* id: %s
* title: [text, a short question title]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !placeholder

[the code below is the starting code in the web editor]
~~~js
// notes on what to return, etc
function doSomething() {
  // return true
}
~~~

##### !end-placeholder

##### !tests

[the mocha tests below will run against the student submission]
~~~js
describe('doSomething', function() {

  it("does what it is supposed to do", function() {
    expect(doSomething(), "Error message").to.deep.eq(true)
  })
})
~~~

##### !end-tests

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const javaTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: code-snippet
* language: java
* id: %s
* title: [text, a short question title]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !setup

[the code below will be added to the beginning of the student submission]
~~~java
// include any imports specific to your tests
import java.io.IOException;

// to allow student to submit simple statements, wrap the submission
//  using the !setup and !tests sections; example below
class VariableChallenge {

    public static String run() {
        // Start Student Code
~~~

##### !end-setup

##### !placeholder

~~~java
[the code below is the starting code in the web editor]
// write code that declares the string foo and sets it to "bar"
// String foo="bar";
~~~

##### !end-placeholder

##### !tests

[the test code below will be added to the end of the student submission]
~~~java
  // End Student Code
  return foo;
  }
}

// public test class name **must** be SnippetTest to match the generated file name
public class SnippetTest {

	@Test
	public void someTest() {
		assertEquals("bar", VariableChallenge.run());
	}
}
~~~

##### !end-tests

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const pythonTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: code-snippet
* language: python3.6
* id: %s
* title: [text, a short question title]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !placeholder

[the code below is the starting code in the web editor]
~~~py
def doSomething():
  '''
  INPUT: 2 dimensional numpy array
  OUTPUT: boolean
  Return true
  '''
#   return 1
~~~

##### !end-placeholder

##### !tests

[the unit tests below will run against the student submission]
~~~py
import unittest
import main as p
import numpy as np

class TestPython1(unittest.TestCase):
  def test_one(self):
    self.assertEqual(1,p.doSomething())
~~~

##### !end-tests

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const sqlTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: code-snippet
* language: sql
* id: %s
* title: [text, a short question title]
* data_path: /[text, the path to the folder with the .sql file]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !placeholder

[the code below is the starting code in the web editor]
~~~sql
-- write a statement to select...
~~~

##### !end-placeholder

##### !tests

[the code below is the sql statement that returns the correct answer]
~~~sql
SELECT these
FROM that
JOIN other
WHERE this
GROUP BY logic
ORDER BY something
~~~

##### !end-tests

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const projectTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: project
* id: %s
* title: [text, a short question title]
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !placeholder

[text, placeholder text for the input field]

##### !end-placeholder

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const testableProjectTemplate = `<!-- >>>>>>>>>>>>>>>>>>>>>> BEGIN CHALLENGE >>>>>>>>>>>>>>>>>>>>>> -->
<!-- Replace everything in square brackets [] and remove brackets  -->

### !challenge

* type: testable-project
* id: %s
* title: [text, a short question title]
* upstream: [URL, the upstream repo URL like https://github.com/gSchool/js-native-array-methods/]
* validate_fork: true
<!-- * points: [1] (optional, the number of points for scoring as a checkpoint) -->
<!-- * topics: [python, pandas] (optional the topics for analyzing points) -->

##### !question

[markdown, your question]

##### !end-question

##### !placeholder

[text, placeholder text for the input field]

##### !end-placeholder

<!-- other optional sections -->
<!-- !hint - !end-hint (markdown, users can see after a failed attempt) -->
<!-- !rubric - !end-rubric (markdown, instructors can see while scoring a checkpoint) -->
<!-- !explanation - !end-explanation (markdown, students can see after answering correctly) -->

### !end-challenge

<!-- ======================= END CHALLENGE ======================= -->`

const configyamlTemplate = `# Config.yaml specifies the content and ordering within a curriculum block repo
#
# Supported Fields
# ==========================
# Standards -- (Standards = Units). An array of Units for a block
# Standard.Title -- The Unit title that shows up on the curriculum overview
# Standard.UID -- A unique ID for the Unit.
# Standard.Description -- The Unit description that shows up on the curriculum overview
# Standard.SuccessCriteria -- An array of success criteria that can be viewed when scoring the checkpoint in a Unit.
# Standard.ContentFiles -- An array of Lessons and (optional) Checkpoint in a Unit.
# Standard.ContentFiles.Type -- 'Lesson' or 'Checkpoint'
# Standard.ContentFiles.UID -- A unique ID for the lesson or checkpoint.
# Standard.ContentFiles.Path -- The absolute path to the Lesson, starting with /.
# Standard.ContentFiles.DefaultVisibility -- (optional) Set to 'hidden' to hide when a course first starts.
# Standard.ContentFiles.Autoscore -- (optional, for Checkpoints only) submit checkpoint scores without review
# Standard.ContentFiles.MaxCheckpointSubmissions -- (optional, for Checkpoints only) limit the number of submissions
# Standard.ContentFiles.TimeLimit -- (optional, for Checkpoints only) the time limit in minutes
#
# Instructions
# ==========================
# Replace everything in square brackets [] and remove brackets
# Add all other Standards, Lessons, and Checkpoints following the pattern below
# All UIDs must be unique within a repo. You can use a uuidgen plugin.

---
Standards:
  - Title: [The unit name]
    UID: [unique-id]
    Description: [The Standard text]
    SuccessCriteria:
      - [The first success criteria]
    ContentFiles:
      - Type: Lesson
        UID: [unique-id]
        Path: /[folder/file.md]
      - Type: Checkpoint
        UID: [unique-id]
        Path: /[folder/file.md]
`

const courseyamlTemplate = `# Course.yaml files specify the grouping and ordering of repos that define a course.
#
# Supported Fields
# ===================
# DefaultUnitVisibility -- (optional) set to 'hidden' to hide all units when a course first starts.
# Course -- The top level array containing the sections of a course
# Course.Section -- An array contining a single array of repos. Content in the same section is grouped together on curriculum homepage.
# Course.Repos --  An array containing block repos that have been published in Learn.
# Course.Repos.URL -- The URL to a block repo that has been published in Learn.
#
# Instructions
# ==========================
# Replace everything in square brackets [] and remove brackets
# Add all other Sections and Repos following the pattern below
# All UIDs must be unique within a repo. You can use a uuidgen plugin.

---
# DefaultUnitVisibility: hidden
Course:
  - Section: [Section name]
    Repos:
      - URL: https://github.com/gSchool/[Repo name]`
