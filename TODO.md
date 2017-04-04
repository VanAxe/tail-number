# tail-number

## TODO

### Dependency Operations

#### MVP

* Read configuration file
    * JSON / YAML
* Read versions
* Compute bump
* Write lock file
    * JSON

#### Change log management

* Read change log
* Collate upstream changes
* Replace bump word for numeral
* Write change log

### Branch operations

#### MVP

* Read stategy
    * Source
    * Destination
    * Operation type
        * Merge / Squash / Rebase
* Read rules and constraints
* Parse commits
    * Authors / Commiter
    * Files changed
    * Commit message
* Merge branch
* Create git tag
* Push to repo

#### GPG Check

* Verify GPG signature authenticity
