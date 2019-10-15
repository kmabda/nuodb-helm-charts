# Publishing an Index

## Initially creating a gh-pages branch

When a project is first setup, the gh-pages branch must be properly configured to enable gh-pages.

To create an empty gh-pages branch:

```bash
git checkout --orphan gh-pages
git add LICENSE .gitignore
git commit -m "initial"
git push --set-upstream origin gh-pages
```

Then select `gh-pages branch` under Github..Settings, Github Pages. Choose the Slate theme.

## Publishing an index

