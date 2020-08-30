define(function(require) {
  var app = require("app");
  var $ = require("jquery");
  var parsers = require("parsers");
  var CurrentUser = require("js/app/apis/user/current");
  return function(vm, cb) {
    var url = app.Host + app.APIList.updateprofile;
    $.post(url, JSON.stringify(vm.Item))
      .done(function(body) {
        var data = parsers.parse200(body);
        CurrentUser(function(){
          cb(data);
        })
      })
      .fail(function(xhr) {
        if (xhr.status === 422) {
          vm.errors = parsers.parse422(xhr.responseJSON);
          cb()
        }
      });
  };
});
