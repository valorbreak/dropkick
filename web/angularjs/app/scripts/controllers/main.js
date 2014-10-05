'use strict';

/**
 * @ngdoc function
 * @name dropkickApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the dropkickApp
 */
angular.module('dropkickApp')
  .controller('MainCtrl', function ($scope) {
    $scope.awesomeThings = [
      'HTML5 Boilerplate',
      'AngularJS',
      'Karma'
    ];
  });
