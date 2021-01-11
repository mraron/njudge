# njudge

Online Judge system written in golang.

## Installation

### 1. Install the dependencies

For example on Ubuntu 20.10:

`$ sudo apt install pandoc g++ fpc golang julia octave python3`

Install [isolate](https://github.com/ioi/isolate) on the system.

Install postgresql on the system and create a database for njudge.

### 2. Clone the repository and create config files

`$ git clone github.com/mraron/njudge.git`

Create `judge.json` and and `web.json` from the appropriate `.example` files, modifying them in respect to your system.

### 3. Migrate database

Run njudge by executing `go build` in the root of the repository and then executing 
`$ ./njudge migrate --up -c web.json`
to migrate the database to the latest version.  

### 4. Run

Run your judger via the command `./njudge judge -c judge.json`
Run the web frontend via the command `./njudge web -c web.json`

In the database or on the admin panel add the judger to the list.


## Modules
* njudge/web: web frontend [![](https://godoc.org/github.com/mraron/njudge/web?status.svg)](http://godoc.org/github.com/mraron/njudge/web)
* njudge/judge: judges programs sent to it by an njudge/web instance [![](https://godoc.org/github.com/mraron/njudge/judge?status.svg)](http://godoc.org/github.com/mraron/njudge/judge)
* njudge/utils/language: sandboxing utilities [![](https://godoc.org/github.com/mraron/njudge/utils/language?status.svg)](http://godoc.org/github.com/mraron/njudge/utils/language)
* njudge/utils/problems: parsing and internal representation of problems [![](https://godoc.org/github.com/mraron/njudge/utils/problems?status.svg)](https://godoc.org/github.com/mraron/njudge/utils/problems)

## Screenshots
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/1.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/2.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/3.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/4.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/5.png" width="25%" height="25%">
<img src="https://raw.githubusercontent.com/mraron/assets/master/njudge/6.png" width="25%" height="25%">

