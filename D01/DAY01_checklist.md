# Introduction

Please respect the following rules:

- Remain polite, courteous, respectful and constructive
throughout the evaluation process. The well-being of the community
depends on it.
- Identify with the person (or the group) evaluated the eventual
dysfunctions of the work. Take the time to discuss
and debate the problems you have identified.
- You must consider that there might be some difference in how your
peers might have understood the project's instructions and the
scope of its functionalities. Always keep an open mind and grade
him/her as honestly as possible. The pedagogy is valid only and
only if peer-evaluation is conducted seriously.
Guidelines
- Only grade the work that is in the student or group's
GiT repository.
- Double-check that the Git repository belongs to the student
or the group. Ensure that the work is for the relevant project
and also check that "git clone" is used in an empty folder.
- Check carefully that no malicious aliases was used to fool you
and make you evaluate something other than the content of the
official repository.
- To avoid any surprises, carefully check that both the evaluating
and the evaluated students have reviewed the possible scripts used
to facilitate the grading.
- If the evaluating student has not completed that particular
project yet, it is mandatory for this student to read the
entire subject prior to starting the defence.
- Use the flags available on this scale to signal an empty repository,
non-functioning program, a norm error, cheating etc. In these cases,
the grading is over and the final grade is 0 (or -42 in case of
cheating). However, with the exception of cheating, you are
encouraged to continue to discuss your work (even if you have not
finished it) in order to identify any issues that may have caused
this failure and avoid repeating the same mistake in the future.
- Remember that for the duration of the defence, no segfault,
no other unexpected, premature, uncontrolled or unexpected
termination of the program, else the final grade is 0. Use the
appropriate flag.

You should never have to edit any file except the configuration file if it exists.
If you want to edit a file, take the time to explicit the reasons with the
evaluated student and make sure both of you are okay with this.

You are allowed to use any of the different tools available on the computer, such as
pprof.

## Mandatory Part

Comparing incomparable

### Common

Project structure
- Check that project only requires running `go build` to produce an executable
- Check that submission includes *go.mod* and *go.sum* (second one should be there in case external dependencies are used)
- Yes No
 
### Exercise 00

Input processing
- Check that application can read provided example XML file without errors
- Check that application can read provided example JSON file without errors
- Try passing JSON/XML files with broken format and check that program prints a meaningful error
Yes No
Reading
- Check that application prints a corresponding JSON when being passed an XML file
- Check that application prints a corresponding XML when being passed an JSON file
- Check that in both cases 4-space indent is used for pretty-printing
Yes No
Interface
- Check that application code contains a `DBReader` interface and two structs that implement this interface - one for XML and one for JSON
- Check that application decides which implementation to use based on specified file extension
Yes No

### Exercise 01

Assessing damage - cakes and ingredients
- Check that adding a new cake or removing old one is correctly detected and printed
- Check that the change in cooking time is detected and printed
- Check that adding or removing an ingredient to an existing cake is correctly detected and printed
- Check that changing the order of ingredients in input JSON/XML files doesn't affect the output
Yes No
Assessing damage - counts and units
- Check that changing units and counts for ingredients is correctly detected and printed
- Check that removal or addition of a unit field is correctly detected and printed
Yes No

### Exercise 02

Afterparty
- Check that at no time both files are fully loaded into memory
- Check that adding or removing new filepath is correctly detected and printed
- Check that the program works correctly with empty files (both old and new)
Yes No

## Ratings

Don't forget to check the flag corresponding to the defense

Ok
Outstanding project
Empty work
Invalid compilation
Crash
Leaks
Wrong result

## Conclusion

Leave a comment on this evaluation