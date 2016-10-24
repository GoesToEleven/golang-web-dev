(function() {
  angular.module("services", []);
  var app = angular.module("app", ["ui.router", "services", "ui.bootstrap"]);

  app.config(function($stateProvider, $urlRouterProvider) {
    $urlRouterProvider.otherwise("/home");

    $stateProvider
      .state("landing", {
        url: "/",
        templateUrl: "res/views/home.html"
      })
      .state("create", {
        url: "/purchaseOrders",
        templateUrl: "res/views/create.html"
      })
      .state("manageSearch", {
        url: "/purchaseOrders/manage/search",
        templateUrl: "res/views/manageSearch.html"
      })
      .state("manage", {
        url: "/purchaseOrders/:poNumber/manage",
        templateUrl: "res/views/manage.html"
      })
      .state("reviewSearch", {
        url: "/purchaseOrders/review/search",
        templateUrl: "res/views/reviewSearch.html"
      })
      .state("review", {
        url: "/purchaseOrders/:poNumber",
        templateUrl: "res/views/review.html"
      })
  });

  app.run(function($state) {
    $state.go("landing");
  });
})()
