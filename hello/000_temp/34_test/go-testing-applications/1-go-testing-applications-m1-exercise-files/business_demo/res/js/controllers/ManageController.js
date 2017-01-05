(function() {

  var ManageController = function(PurchaseOrderService, $stateParams) {
    var self = this;
    this.poNumber = $stateParams.poNumber;

    this.doSearch = function(e) {
      if (e.keyCode == 13) { //enter
          location.hash="#/purchaseOrders/" + this.poNumber + "/manage";
      }
    }

    this.purchaseOrder = null;
    this.totalCost = 0;

    PurchaseOrderService.getOrder(this.poNumber).then(function(order){
      self.purchaseOrder = order;
      self.totalCost = self.purchaseOrder.totalCost();
    });
  }

  angular.module("app").controller("ManageController", ManageController);
})();
