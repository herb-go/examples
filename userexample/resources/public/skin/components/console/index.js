define(function(require) {
  return {
    name: "console",
    template: require("text!./index.html"),
    props: ["user"],
    data: function() {
      return {};
    }
  };
});
