#!/bin/sh

AWS_PROFILE="$AWS_PROFILE" go run .

selected_profile="$(cat ~/.awsp)"

if [ -z "$selected_profile" ]
then
  unset AWS_PROFILE
else
  export AWS_PROFILE="$selected_profile"
fi