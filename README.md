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

Then under the project, edit the project title, setting the Website as `https://nuodb.github.io/nuodb-helm-charts/`.

## Publishing an index

On the master branch:

```bash
./scripts/index
git checkout gh-pages
git add index.html
git commit -m "update index"
git push
```
