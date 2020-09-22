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

Project structure
- Check that project only requires running `go build` to produce an executable
- Check that submission includes *go.mod* and *go.sum* (second one should be there in case external dependencies are used)
- Yes No

### Exercise 00

Basic functionality
- Check that application finds recursively and prints files, directories and symlinks if given a directory which has those
Yes No
Finding files
- Check that application finds and prints only files if given a '-f' option
Yes No
Finding files with an extension
- Check that application finds and prints only files only files with a given extension if given '-f' and '-ext' options
- Check that '-ext' option is only applied when '-f' is specified
Yes No
Finding directories
- Check that application finds and prints only directories if given a '-d' option
Yes No
Finding symlinks
- Check that application finds and prints only symlinks if given a '-sl' option
Yes No
Printing symlinks
- Check that application successfully resolves symlinks and prints both link and original paths for them, separated by '->'
Yes No
Broken symlinks
- Check that application prints both link and '[broken]' for broken symlinks
Yes No
- Check that application works correctly with two or three options from ['-f', '-d', '-sl'] specified simultaneously
Yes No
- Check that application prints a meaningful error message when no directory path is given or when it doesn't exist
Yes No

### Exercise 01

- Check that application calculates a number of lines in one file correctly
Yes No
- Check that application calculates a number of words in one file correctly
Yes No
- Check that application calculates a number of characters in one file correctly
Yes No
- Check that application calculates all three parameters for multiple files simultaneously
Yes No
- Check that if multiple files are specified, they are processed concurrently in goroutines
Yes No
- Check that application's output is separated by tab symbol
Yes No

### Exercise 02

- Check that application treats all arguments as a command to run
Yes No
- Check that application doesn't die with runtime error when given non-existing or invalid command
Yes No

### Exercise 03

- Check that application creates a proper tar.gz archive with a file inside
Yes No
- Check that application correctly determines original file's mtime
Yes No
- Check that if multiple files are specified, they are processed concurrently in goroutines
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