# golife

This is a simple webserver that will serve JSON over a REST API.

More details on the purpose of the project will be added later in the course of this project.

## Run tests

From the root directory of the project simply run

  `ginkgo -r`

## Deploy to Cloudfoundry

Essentially, follow the instructions mentioned [here](http://mmcgrana.github.io/2012/09/getting-started-with-go-on-heroku)
Cloudfoundry has the same deployment process as Heroku and it has a go buildpack by default. This can be seen by running the command `cf buildpacks`

Also the application reads an environment variable `PORT`. This port is allocated by the DEA. More information can be found [here](http://docs.run.pivotal.io/devguide/deploy-apps/environment-variable.html#PORT)
