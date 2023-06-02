# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[unreleased]: https://github.com/mraron/njudge/compare/v0.1.0...HEAD
[0.0.1]: https://github.com/mraron/njudge/releases/tag/v0.1.0