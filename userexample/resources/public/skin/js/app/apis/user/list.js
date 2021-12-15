define(function(require) {
    var app = require("app");
    var $ = require("jquery");
    var parsers = require("parsers");
    return function(vm, cb) {
      var url = app.Host + app.APIList.list;
      var data={
        last:vm.Last,
        rev:vm.Rev?"true":"",
      }
      $.get(url, data)
        .done(function(body) {
          var data = parsers.parse200(body);
          vm.Items=parsers.parseItems(body);
          vm.Iter=body.Iter;
          cb(data);
        })
        .fail(function(xhr) {});
    };
  });
  