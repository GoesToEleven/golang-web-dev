(function() {
  var receivers, vendors;

  var ShippingService = function($q, $http) {
    this.getReceivers = function() {
      return $q(function(resolve, reject) {
        if (receivers) {
          resolve(receivers);
        } else {
          $http.get("/api/receivers").then(function(response){
            receivers = response.data;
            resolve(receivers);
          });
        }
      });
    }

    this.getVendors = function() {
      return $q(function(resolve, reject) {
        if (vendors) {
          resolve(vendors);
        } else {
          $http.get("/api/vendors").then(function(response) {
            vendors = response.data;

            resolve(vendors);
          });
        }
      });
    }
  }

  angular.module("services").service("ShippingService", ShippingService);
})()
