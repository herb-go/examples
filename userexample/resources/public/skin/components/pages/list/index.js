define(function(require) {
  var itemlist = require("js/app/apis/user/list");
  var itemenable = require("js/app/apis/user/enable");
  var itemdisable = require("js/app/apis/user/disable");

  var lodash = require("lodash");
  return {
    name: "compomentname",
    watch: {
      $route: function(to, from) {
        this.load();
      }
    },
    mounted: function() {
      this.load();
    },
    methods: {
      load: function() {
        var self = this;
        page = this.$route.query.page;
        if (!page) {
          page = 1;
        }
        this.CurrentPage = page * 1;
        this.Sort = this.$route.query.sort;
        this.Asc = this.$route.query.order == "ascending";
        itemlist(self, function() {});
      },
      onPage: function(page) {
        this.CurrentPage = page;
        var query = lodash.clone(this.$route.query);
        query.page = page;
        this.$router.push({ query: query });
      },
      handleEnable: function(item) {
        var self = this;
        self.errors = [];
        this.$confirm("此操作将激活用户, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(function() {
            itemenable(self, item.ID, function() {
              self.load();
            });
          })
          .catch(function() {});
      },
      handleDisable: function(item) {
        var self = this;
        self.errors = [];
        this.$confirm("此操作将禁用用户, 是否继续?", "提示", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        })
          .then(function() {
            itemdisable(self, item.ID, function() {
              self.load();
            });
          })
          .catch(function() {});
      },
      onSort: function(data) {
        this.Sort = data.prop;
        this.Asc = data.order == "ascending";
        var query = lodash.clone(this.$route.query);
        query.sort = data.prop;
        query.order = data.order;
        this.$router.push({ query: query });
      }
    },
    template: require("text!./index.html"),
    data: function() {
      return {
        Items: [],
        Sort: "",
        Asc: true,
        errors: [],
        Count: null,
        CurrentPage: 1
      };
    }
  };
});
