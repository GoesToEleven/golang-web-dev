(function() {
  var PurchaseOrderService = function($q, $http) {
    this.createOrder =  function(order) {
      return $q(function(resolve, reject) {
        order.details.forEach(function(d) {
          d.pricePer = parseFloat(d.pricePer);
          d.quantity = parseInt(d.quantity);
        })
        $http.post("/api/purchaseOrders", order).then(function(response) {
            resolve(response.data);
        });
      });
    }

    this.getOrder = function(poNumber) {
      return $q(function(resolve, reject) {
        $http.get("/api/purchaseOrders/"+poNumber).then(function(response){
          var result = new PurchaseOrder(response.data);
          result.details = response.data.details.map(function(d) {
            return new PODetail(d);
          });
          resolve(result);
        });
      });
    }

  }


  angular.module("services").service("PurchaseOrderService", PurchaseOrderService);
})();
