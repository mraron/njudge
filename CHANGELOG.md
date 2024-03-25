# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

# Fixed
- User registration (8a00521)
- Sorting was ignored for submission lists (6833e08)
- Loading tags for problems, category filter in problemlist, solver_count query (63f405a)
- Admin panel (661484b, 1ee0070)
- Syntax of migration 10 (e35d4cc)
- Category link in problem sets (d9a5c96)
- Subtask scoring (6fee8f0)
- Fix bug related to long lines and cms whitediff checker (7289bb8)

# Added
- Tests to validate task_yaml ScoringMin behaviour (67a068b)
- Problem and category visibility (5452ee8)
- Optimized submission queries for glue (888a0ae)
- Cached solved status queries (18c5177)

# Changed
- bootstrap-icons (755b9cf)
- Dependencies (74e7c2d, 7c3199e, a0bc370)

## [0.3.1] - 2023-11-28

### Fixed
- Point calculation (de2492f)
- Fix admin panel (ef9d84e)

## [0.3.0] - 2023-11-27

### Added
- Functionality to reset passwords (e5b4b2)

### Fixed
- Crash in Codeforces Feedback problems (6782c6)
- Activating email (text<->html) (41dac0)

## Changed
- Updated to sqlboiler v4.15 and go 1.21
- Database was refactored 
- Business logic was refactored (added in-memory persistance)

The previous two changes are from 65ec4f to 19c523

## [0.2.0] - 2023-06-14

### Added
- logging in case of callback failure in judge
- Auto compilation for cpp checkers in problem.yaml config type
- Support for outputonly tasks in polygon and problem_yaml config types (2cdd17)
- Support for empty task_type_parameters in task_yaml (defaulting to sum subtask with evenly distributed points) (e5f5a4)

### Changed
- testlib checker to only support quitp type partial scoring
- polygon config to not generate html by default from the problem-properties.json
- toString template func to be smarter
- Moved MemoryLimit and TimeLimit to the Problem interface, since they're not used in the evaluation process: only the status skeleton matters.
- Optimize judge Dockerfile (5347e1)
- In outputonly tasktype if the file is not found in the zip, then the verdict is now VerdictDR (2cdd17)
- task_yaml: properly set locale from primary_language (033f1a)
- Optimized stderr capturing in isolate sandbox and added tests (bc0560)

### Fixed
- Updating language and problem list every 20 second.
- Task archive only displaying the toplevel categories.
- Tr template.Funcs's arguments
- testset.FirstNonAC
- problem_yaml: set memory and time limit correctly
- Clear filter button on the problemset list page (eb7896)
- Judge didn't spawn new gorutines for the workers (923f67) 

## [0.1.0] - 2023-06-02

### Added

- language.Store interface.
- Error reporting to the Status returned by the judge.
- ACE editor for submitting on the problem page.
- Support for regexes (list of testcases used to simulate dependencies) in task_yaml
- For the previous point, caching of testcases in batch tasktype if their InputPath points to the same location.
- Flash messages (messages stored in cookies, useful to persist message after a redirect)
- cpp.AutoCompile which automatically extracts headers and performs an unsafe compilation of task materials
- language.StoreAllExcept: convenience method to filter out a list of language ids (mostly zip)
- CSRF protection to all forms.
- New favicons.
- English translation.
- New mail template.
- Embedding in case of production and auto-reloading of templates in case of development mode.
- Profile settings page: a way for users to change their password and set the visibility of tags for unsolved problems.
- Context to a bunch of places.
- Redirection in case using helpers.LoginRequired, it's now redirecting to /user/login?next={url} and /user/login will redirect back to url.

### Changed

- To use afero.FS in problem configs.
- Refactored judge service to be more robust.
- Now internal errors are displayed to user.
- In task_yaml if GEN is present but also there's a `score_type_parameters` field in the yaml, prefer the latter (for regexes)
- Renamed problems.ConfigStore to problems.ConfigList
- Some mapstructure annotations of the web configuration structs (#83)
- Back the migration dependency to the original github.com/golang-migrate/migrate, since they now also support go1.18
- Moved out a bunch of business logic and data modeling from the handlers into the new services and domain packages.
- Template functions that require context, now they're injected automatically.
- language.Verdict's to idiomatic casing
- problem_json to problem_yaml

### Removed

- HTMLStatements and PDFStatements, it's now recommended to use the filtering methods of problems.Contents
- Score, MaxScore, Verdict, IsAC, MaxMemoryUsage, MaxTimeSpent, FirstNonAC, IndexTestcase of Status. Now most of these functionality is in Testsets.

### Fixed

- Crash when admin panel is visited without logging in
- Workflow

[unreleased]: https://github.com/mraron/njudge/compare/v0.3.1...HEAD
[0.3.1]: https://github.com/mraron/njudge/releases/tag/v0.3.1
[0.3.0]: https://github.com/mraron/njudge/releases/tag/v0.3.0
[0.2.0]: https://github.com/mraron/njudge/releases/tag/v0.2.0
[0.1.0]: https://github.com/mraron/njudge/releases/tag/v0.1.0
