define(function(require) {
    var app = require("app");
    var $ = require("jquery");
    var parsers = require("parsers");
    return function(vm, cb) {
      var url = app.Host + app.APIList.actives;
      $.get(url, {page:vm.CurrentPage})
        .done(function(body) {
          var data = parsers.parse200(body);
          vm.Items=parsers.parseItems(body);
          cb(data);
        })
        .fail(function(xhr) {});
    };
  });
  