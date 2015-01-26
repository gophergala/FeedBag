angular.module "ui"
  .directive "eventTemplate", eventTemplate

eventTemplate = ->
  directive =
    restrict: 'E'
    scope:
      data: "="
    templateUrl: template
    link: link

  directive

  template = (el, attr)->
    attr.templateUrl

  link = (scope, el) ->

