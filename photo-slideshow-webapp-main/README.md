# moje notatki

## backendServer:

server: ma glownego maina

photodb: ma tutaj baze danych sqlite, jak nie ma bazy to tworzy nowa

photoprocessing: to olej, to zmienia z heic(iphonowy format) na jpegi i ogarnia orientacje zdjec

vendor tez olej, to sa moduly pobrane calkowicie

## frontend

olej w sumie, zrobisz w reactie ostatecznie, najwyzej jakies funkcje podpatrz

## imageAutomation - olej, nie uzyjesz tego

## dockerfile - wazne, przejrzyj 

## to sa jakies dziwne templatey do "helm" to chyba cos do deploymentu, nie do konca widze dlaczego on to uzywa i po co


























# Full Stack Social Slideshow Photo Sharing App

This project contains all the code necessary to deploy a photo sharing party slideshow app.

Features:

- REST API written in Go for posting photos, logging in, and retrieving photos to display on a big screen TV (see backendServer directory)
- Front end single page application written in VueJS. The front end contains three views: The login, the photo upload view and the big-screen display slideshow (see frontend directory)
- Lightweight application that designed for running on a laptop that monitors a directory for photos and uploads them to the API. (see imageAutomation)
  - This was designed for the specific purpose of a photo booth where a fixed camera was taking photos and putting them on a laptop. This is completely optional. The app works standalone with smartphones
- Dockerfile and Helm chart included for deployment to kubernetes

## Watch the video series on the making of this project!

<a href="https://www.youtube.com/playlist?list=PLSvCAHoiHC_r0zdPt37-JfG85WBGD1cey" target="_blank">
<img src="http://img.youtube.com/vi/wqcg9w_Q0iA/mqdefault.jpg" alt="Watch the series" width="240" height="180" border="10" />
</a>

## How to install

1. watch the series
2. clone the project
3. use the helm chart to deploy after making your necessary changes

## Find a Bug?

Please submit an issue or a pull request. Thanks!

## Like this project?

If you are feeling generous, buy me a coffee! - https://www.buymeacoffee.com/askcloudtech
