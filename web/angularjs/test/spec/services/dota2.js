'use strict';

describe('Service: dota2', function () {

  // load the service's module
  beforeEach(module('dropkickApp'));

  // instantiate service
  var dota2;
  beforeEach(inject(function (_dota2_) {
    dota2 = _dota2_;
  }));

  it('should do something', function () {
    expect(!!dota2).toBe(true);
  });

});
