# AWSP - AWS Profile Switcher

> a go clone of [johnnyopao/awsp](https://github.com/johnnyopao/awsp)

Easily switch between AWS Profiles

<img src="https://github.com/johnnyopao/awsp/raw/master/demo.gif" width="500">

## Prereqs
Setup your profiles using the aws cli

```sh
aws configure --profile PROFILE_NAME
```

You can also leave out the `--profile PROFILE_NAME` param to set your `default` credentials

Refer to this doc for more information
https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html


## Homebrew Install

`brew install electblake/awsp/awsp`

Or `brew tap electblake/awsp` and then `brew install awsp`.

### Updates

Highly reccomend tap first `brew tap electblake/awsp`

and then simple `brew upgrade awsp`

## Usage
```sh
awsp
```

## Show your AWS Profile in your shell prompt
For better visibility into what your shell is set to it's helpful to configure your prompt to show the value of the env variable `AWS_PROFILE`.

<img src="https://github.com/johnnyopao/awsp/raw/master/screenshot.png" width="300">

Here's a sample of my zsh prompt config using oh-my-zsh themes

```sh
function aws_prof {
  local profile="${AWS_PROFILE:=default}"

  echo "%{$FG[238]%}aws(%{$FG[028]%}${profile}%{$FG[238]%})%{$reset_color%} "
}
```

```sh
PROMPT='OTHER_PROMPT_STUFF $(aws_prof)'
```


## References and Inspiration
---
- https://github.com/johnnyopao/awsp
- https://github.com/99designs/aws-vault
- https://github.com/kellyp/lowprofile