# frozen_string_literal: true
# This file is @generated
require "json"

module Svix
  class EnvironmentOut
    attr_accessor :connectors
    attr_accessor :created_at
    attr_accessor :event_types
    attr_accessor :settings
    attr_accessor :version

    ALL_FIELD ||= ["connectors", "created_at", "event_types", "settings", "version"].freeze
    private_constant :ALL_FIELD

    def initialize(attributes = {})
      unless attributes.is_a?(Hash)
        fail(ArgumentError, "The input argument (attributes) must be a hash in `Svix::EnvironmentOut` new method")
      end

      attributes.each do |k, v|
        unless ALL_FIELD.include?(k.to_s)
          fail(ArgumentError, "The field #{k} is not part of Svix::EnvironmentOut")
        end

        instance_variable_set("@#{k}", v)
        instance_variable_set("@__#{k}_is_defined", true)
      end
    end

    def self.deserialize(attributes = {})
      attributes = attributes.transform_keys(&:to_s)
      attrs = Hash.new
      attrs["connectors"] = attributes["connectors"].map { |v| Svix::ConnectorOut.deserialize(v) }
      attrs["created_at"] = DateTime.rfc3339(attributes["createdAt"]).to_time
      attrs["event_types"] = attributes["eventTypes"].map { |v| Svix::EventTypeOut.deserialize(v) }
      attrs["settings"] = attributes["settings"]
      attrs["version"] = attributes["version"]
      new(attrs)
    end

    def serialize
      out = Hash.new
      out["connectors"] = @connectors.map { |v| v.serialize } if @connectors
      out["createdAt"] = Svix::serialize_primitive(@created_at) if @created_at
      out["eventTypes"] = @event_types.map { |v| v.serialize } if @event_types
      out["settings"] = Svix::serialize_primitive(@settings) if @settings
      out["version"] = Svix::serialize_primitive(@version) if @version
      out
    end

    # Serializes the object to a json string
    # @return String
    def to_json
      JSON.dump(serialize)
    end
  end
end
