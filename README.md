# Introduction

This is the common package using reusable code for [e-commerce-microservice repo](https://github.com/Kiyosh31/e-commerce-microservice)
here I have all the code reused alongside the micro services

# Instructions

To publish new versions of this package do the next steps

1. Make a modification, commit and push changes

```console
git add .
git commit -m "Message"
git push
```

1. Generate new tag

```console
# git tag "<version>"
git tag <v1.0.0>
```

2. Push to GH

```console
git push --tags
```

3. Install the package in your project

```console
go get github.com/Kiyosh31/e-commerce-microservice-common
```
