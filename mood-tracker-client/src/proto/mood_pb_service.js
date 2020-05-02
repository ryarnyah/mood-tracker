// package: 
// file: mood.proto

var mood_pb = require("./mood_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Mood = (function () {
  function Mood() {}
  Mood.serviceName = "Mood";
  return Mood;
}());

Mood.AddEntry = {
  methodName: "AddEntry",
  service: Mood,
  requestStream: false,
  responseStream: false,
  requestType: mood_pb.AddEntryRequest,
  responseType: mood_pb.AddEntryResponse
};

Mood.GetMoodFromEntry = {
  methodName: "GetMoodFromEntry",
  service: Mood,
  requestStream: false,
  responseStream: false,
  requestType: mood_pb.GetMoodFromEntryRequest,
  responseType: mood_pb.GetMoodFromEntryResponse
};

Mood.GetMood = {
  methodName: "GetMood",
  service: Mood,
  requestStream: false,
  responseStream: false,
  requestType: mood_pb.GetMoodRequest,
  responseType: mood_pb.GetMoodResponse
};

Mood.CreateMood = {
  methodName: "CreateMood",
  service: Mood,
  requestStream: false,
  responseStream: false,
  requestType: mood_pb.CreateMoodRequest,
  responseType: mood_pb.CreateMoodResponse
};

exports.Mood = Mood;

function MoodClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

MoodClient.prototype.addEntry = function addEntry(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Mood.AddEntry, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MoodClient.prototype.getMoodFromEntry = function getMoodFromEntry(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Mood.GetMoodFromEntry, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MoodClient.prototype.getMood = function getMood(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Mood.GetMood, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

MoodClient.prototype.createMood = function createMood(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Mood.CreateMood, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.MoodClient = MoodClient;

