define(function(require) {
  var app = {};
  app.Host = "";
  app.APIList = {
    current: "/api/current",
    logout: "/api/logout",
    login: "/api/login",
    register: "/api/register",
    actives:"/api/actives",
    list:"/api/list",
    revoke:"/api/revoke",
    updatepassword:"/api/updatepassword"
  };
  return app;
});
