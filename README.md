# Gitlab Registry Cleaner

Gitlab Registry Cleaner is a tiny container which can be used to remove images from gitlab registry.

## Usage

1. Create gitlab access token. Read how to [here](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#creating-a-personal-access-token).

1. Run container in your `.gitlab-ci.yml`
```sh
docker run \
 && -e ACCESS_TOKEN='<your_gitlab_access_token>' \
 && -e PROJECT_ID='$CI_PROJECT_ID' \
 && -e IMAGE_TAG='$CI_BUILD_REF_NAME' \
 && -e IMAGE_LOCATION='$GITLAB_REGISTRY/$CI_PROJECT_PATH/<your_path>' \
 && devsgroup/gitlab-registry-cleaner:latest
```

E.g.
```yml
delete_frontend_image:
  stage: clean
  script:
    - |
      docker run \
        -e ACCESS_TOKEN='$GITLAB_ACCESS_TOKEN' \
        -e PROJECT_ID='$CI_PROJECT_ID' \
        -e IMAGE_TAG='$CI_BUILD_REF_NAME' \
        -e IMAGE_LOCATION='$GITLAB_REGISTRY/$CI_PROJECT_PATH/web_frontend' \
        devsgroup/gitlab-registry-cleaner:latest
  only:
    - branches
  except:
    - master
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)