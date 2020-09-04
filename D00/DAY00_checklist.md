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

Statistics being handy
Project structure
- Check that project only requires running `go build` to produce an executable
- Check that submission includes *go.mod* and *go.sum* (second one should be there in case external dependencies are used)
- Yes No
Input processing
- Check that application works when being passed a sequence of integer numbers, separated by newlines
- Try cases when input is sorted and when it isn't
- Try special cases like an empty string with newline, value out of [-100000, 100000] bounds, characters
- Check that application can be run in such way it prints only a specified subset of metrics.
Yes No
Mean calculation
- Check that mean is calculated correctly for both odd and even number of inputs
- Check that the output is rounded to two decimal points
Yes No
Median calculation
- Check that median is calculated correctly for both odd and even number of inputs
- Check that the output is rounded to two decimal points
Yes No
Mode calculation
- Check that mode is always equal to the most frequent number in the input
- Check that if there are multiple most frequent numbers, then mode is equal to the smallest one among those
- Check that the output is rounded to two decimal points
Yes No
SD calculation
- Check that SD calculation works even if mean calculation is disabled
- Check that SD is calculated correctly, being equal to either regular SD or a population SD
- Check that the output is rounded to two decimal points
Yes No

## Ratings

Donâ€™t forget to check the flag corresponding to the defense

Ok
Outstanding project
Empty work
Invalid compilation
Crash
Leaks
Wrong result

## Conclusion

Leave a comment on this evaluation