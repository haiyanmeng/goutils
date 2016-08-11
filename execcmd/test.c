#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/types.h>
#include <sys/wait.h>

/* This program forks and and the prints whether the process is
 *   - the child (the return value of fork() is 0), or
 *   - the parent (the return value of fork() is not zero)
 *
 * When this was run 100 times on the computer the author is
 * on, only twice did the parent process execute before the
 * child process executed.
 *
 * Note, if you juxtapose two strings, the compiler automatically
 * concatenates the two, e.g., "Hello " "world!"
 */

int main( void ) {
	int pid = fork();

	if ( pid == 0 ) {
		printf( "This is being printed from the child process %ld\n",  (long)getpid());
		char *binary = "/usr/local/sbin/runc";
		char *cmd[] = { "/usr/local/sbin/runc", "create", "--bundle", "/home/hmeng/go/src/github.com/hmeng-19/myrunc/c4", "1", (char *)0 };
		execvp(binary, cmd);
		_exit(EXIT_FAILURE);
	} else {
		printf( "This is being printed in the parent process:\n"
		        " - the process identifier (pid) of the child is %d\n", pid );
		int status;
		waitpid(pid, &status, 0);
		fprintf(stdout, "the exit status of the child process is %d\n", status);
		
	}

	return 0;
}
