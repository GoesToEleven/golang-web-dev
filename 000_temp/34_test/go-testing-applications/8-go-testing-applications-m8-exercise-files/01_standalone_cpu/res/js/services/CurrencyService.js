(function() {
  var currencies = null;

  var CurrencyService = function($q, $http) {
    this.getCurrencies = function() {
        return $q(function(resolve, reject) {
          if (currencies) {
            resolve(currencies);
          } else {
            $http.get("/api/currencies").then(function(response){
              currencies = response.data;

              resolve(currencies);
            });
        }
      });
    }
  };

  angular.module("services").service("CurrencyService", CurrencyService);
})();
