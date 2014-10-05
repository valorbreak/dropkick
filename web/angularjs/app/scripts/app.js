'use strict';

/**
 * @ngdoc overview
 * @name dropkickApp
 * @description
 * # dropkickApp
 *
 * Main module of the application.
 */
var dependencies = [
  'ngAnimate',
  'ngCookies',
  'ngResource',
  'ngRoute',
  'ngSanitize',
  'ngTouch'
]

var app = angular.module('dropkickApp', dependencies);

var routeConfig = function ($routeProvider) {
  $routeProvider
    .when('/', {
      templateUrl: 'views/main.html',
      controller: 'MainCtrl'
    })
    .when('/about', {
      templateUrl: 'views/about.html',
      controller: 'AboutCtrl'
    })
    .when('/contact', {
      templateUrl: 'views/contact.html',
      controller: 'AboutCtrl'
    })
    .when('/events', {
      templateUrl: 'views/events.html',
      controller: 'EventsCtrl'
    })
    .when('/dota', {
      templateUrl: 'views/redditdota.html',
      controller: 'RedditdotaCtrl'
    })
    .otherwise({
      redirectTo: '/'
    });
}

app.config(routeConfig);
