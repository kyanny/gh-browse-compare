# gh-browse-compare

A `gh` extension to aim to be a replacement of `hub compare`.

```
$ gh browse-compare v1.0.0..main
$ gh browse-compare v1.0.0...main
```

You may want to add an alias.

```
$ gh alias set compare 'browse-compare'
```

This extension is an experimental implementation, so that it does not validate the command line argument and error handling is rough.

I created this extension rather than directly contributing to `cli/cli` to close https://github.com/cli/cli/issues/4433 because I need to comply with [this guidance](https://github.com/cli/cli/blob/trunk/docs/working-with-us.md). If I get a positive feedback, I am happy to submit a pull request to `cli/cli`.
