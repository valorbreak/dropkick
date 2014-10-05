'use strict';

describe('Controller: RedditdotaCtrl', function () {

  // load the controller's module
  beforeEach(module('dropkickApp'));

  var RedditdotaCtrl,
    scope;

  // Initialize the controller and a mock scope
  beforeEach(inject(function ($controller, $rootScope) {
    scope = $rootScope.$new();
    RedditdotaCtrl = $controller('RedditdotaCtrl', {
      $scope: scope
    });
  }));

  it('should attach a list of awesomeThings to the scope', function () {
    expect(scope.awesomeThings.length).toBe(3);
  });
});
