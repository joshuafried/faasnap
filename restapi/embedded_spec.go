// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "FaaSnap API",
    "title": "faasnap",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/functions": {
      "get": {
        "description": "Return a list of functions",
        "responses": {
          "200": {
            "description": "List of functions",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Function"
              }
            }
          }
        }
      },
      "post": {
        "description": "Create a new function",
        "parameters": [
          {
            "name": "function",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/invocations": {
      "post": {
        "description": "Post an invocation",
        "parameters": [
          {
            "name": "invocation",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Invocation"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "duration": {
                  "type": "number"
                },
                "result": {
                  "type": "string"
                },
                "traceId": {
                  "type": "string"
                },
                "vmId": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/metrics": {
      "get": {
        "description": "Metrics",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/net-ifaces/{namespace}": {
      "put": {
        "description": "Put a vm network",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "name": "interface",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "guest_addr": {
                  "type": "string"
                },
                "guest_mac": {
                  "type": "string"
                },
                "host_dev_name": {
                  "type": "string"
                },
                "iface_id": {
                  "type": "string"
                },
                "unique_addr": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/snapshots": {
      "put": {
        "description": "Put snapshot (copy)",
        "parameters": [
          {
            "type": "string",
            "name": "from_snapshot",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "mem_file_path",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Snapshot"
            }
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      },
      "post": {
        "description": "Take a snapshot",
        "parameters": [
          {
            "name": "snapshot",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Snapshot"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Snapshot"
            }
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/snapshots/{ssId}": {
      "patch": {
        "description": "Change snapshot state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "name": "state",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "dig_hole": {
                  "type": "boolean"
                },
                "drop_cache": {
                  "type": "boolean"
                },
                "load_cache": {
                  "type": "boolean"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/snapshots/{ssId}/mincore": {
      "get": {
        "description": "Get mincore state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Mincore state",
            "schema": {
              "type": "object",
              "properties": {
                "n_nz_regions": {
                  "type": "integer"
                },
                "n_ws_regions": {
                  "type": "integer"
                },
                "nlayers": {
                  "type": "integer"
                },
                "nz_region_size": {
                  "type": "integer"
                },
                "ws_region_size": {
                  "type": "integer"
                }
              }
            }
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      },
      "put": {
        "description": "Put mincore state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "source",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      },
      "post": {
        "description": "Add mincore layer",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "name": "layer",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "fromDiff": {
                  "type": "string"
                },
                "position": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      },
      "patch": {
        "description": "Change mincore state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "name": "state",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "drop_ws_cache": {
                  "type": "boolean"
                },
                "from_records_size": {
                  "type": "integer"
                },
                "inactive_ws": {
                  "type": "boolean"
                },
                "interval_threshold": {
                  "type": "integer"
                },
                "mincore_cache": {
                  "type": "array",
                  "items": {
                    "type": "integer"
                  }
                },
                "size_threshold": {
                  "type": "integer"
                },
                "to_ws_file": {
                  "type": "string"
                },
                "trim_regions": {
                  "type": "boolean"
                },
                "zero_ws": {
                  "type": "boolean"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/snapshots/{ssId}/reap": {
      "get": {
        "description": "get reap state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      },
      "delete": {
        "description": "delete reap state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      },
      "patch": {
        "description": "Change reap state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "name": "cache",
            "in": "body",
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/ui": {
      "get": {
        "description": "UI",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/ui/data": {
      "get": {
        "description": "UI",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/vmms": {
      "post": {
        "description": "Create a VMM in the pool",
        "parameters": [
          {
            "name": "VMM",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "enableReap": {
                  "type": "boolean"
                },
                "namespace": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VM"
            }
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/vms": {
      "get": {
        "description": "Returns a list of active VMs",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/VM"
              }
            }
          }
        }
      },
      "post": {
        "description": "Create a new VM",
        "parameters": [
          {
            "name": "VM",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "func_name": {
                  "type": "string"
                },
                "namespace": {
                  "type": "string"
                },
                "ssId": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VM"
            }
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    },
    "/vms/{vmId}": {
      "get": {
        "description": "Describe a VM",
        "parameters": [
          {
            "type": "string",
            "name": "vmId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VM"
            }
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      },
      "delete": {
        "description": "Stop a VM",
        "parameters": [
          {
            "type": "string",
            "name": "vmId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "$ref": "#/responses/400Error"
          }
        }
      }
    }
  },
  "definitions": {
    "Function": {
      "type": "object",
      "required": [
        "func_name"
      ],
      "properties": {
        "func_name": {
          "type": "string"
        },
        "image": {
          "type": "string"
        },
        "kernel": {
          "type": "string"
        },
        "mem_size": {
          "type": "integer"
        },
        "vcpu": {
          "type": "integer"
        }
      }
    },
    "Invocation": {
      "type": "object",
      "required": [
        "func_name"
      ],
      "properties": {
        "enableReap": {
          "type": "boolean"
        },
        "func_name": {
          "type": "string"
        },
        "loadMincore": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "mincore": {
          "type": "integer",
          "default": -1
        },
        "mincore_size": {
          "type": "integer"
        },
        "namespace": {
          "type": "string"
        },
        "overlay_regions": {
          "type": "boolean"
        },
        "params": {
          "type": "string"
        },
        "ssId": {
          "type": "string"
        },
        "use_mem_file": {
          "type": "boolean"
        },
        "use_ws_file": {
          "type": "boolean"
        },
        "vmId": {
          "type": "string"
        },
        "vmm_load_ws": {
          "type": "boolean"
        },
        "wsFileDirectIo": {
          "type": "boolean"
        },
        "wsSingleRead": {
          "type": "boolean"
        }
      }
    },
    "Snapshot": {
      "type": "object",
      "required": [
        "vmId"
      ],
      "properties": {
        "interval_threshold": {
          "type": "integer"
        },
        "mem_file_path": {
          "type": "string"
        },
        "record_regions": {
          "type": "boolean"
        },
        "size_threshold": {
          "type": "integer"
        },
        "snapshot_path": {
          "type": "string"
        },
        "snapshot_type": {
          "type": "string"
        },
        "ssId": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "vmId": {
          "type": "string"
        }
      }
    },
    "VM": {
      "type": "object",
      "required": [
        "vmId"
      ],
      "properties": {
        "state": {
          "type": "string"
        },
        "vmConf": {
          "type": "object"
        },
        "vmId": {
          "type": "string"
        },
        "vmPath": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "400Error": {
      "description": "Invalid request",
      "schema": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "FaaSnap API",
    "title": "faasnap",
    "version": "1.0.0"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/functions": {
      "get": {
        "description": "Return a list of functions",
        "responses": {
          "200": {
            "description": "List of functions",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Function"
              }
            }
          }
        }
      },
      "post": {
        "description": "Create a new function",
        "parameters": [
          {
            "name": "function",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/invocations": {
      "post": {
        "description": "Post an invocation",
        "parameters": [
          {
            "name": "invocation",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Invocation"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "type": "object",
              "properties": {
                "duration": {
                  "type": "number"
                },
                "result": {
                  "type": "string"
                },
                "traceId": {
                  "type": "string"
                },
                "vmId": {
                  "type": "string"
                }
              }
            }
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/metrics": {
      "get": {
        "description": "Metrics",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/net-ifaces/{namespace}": {
      "put": {
        "description": "Put a vm network",
        "parameters": [
          {
            "type": "string",
            "name": "namespace",
            "in": "path",
            "required": true
          },
          {
            "name": "interface",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "guest_addr": {
                  "type": "string"
                },
                "guest_mac": {
                  "type": "string"
                },
                "host_dev_name": {
                  "type": "string"
                },
                "iface_id": {
                  "type": "string"
                },
                "unique_addr": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/snapshots": {
      "put": {
        "description": "Put snapshot (copy)",
        "parameters": [
          {
            "type": "string",
            "name": "from_snapshot",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "mem_file_path",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/Snapshot"
            }
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      },
      "post": {
        "description": "Take a snapshot",
        "parameters": [
          {
            "name": "snapshot",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Snapshot"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Success",
            "schema": {
              "$ref": "#/definitions/Snapshot"
            }
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/snapshots/{ssId}": {
      "patch": {
        "description": "Change snapshot state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "name": "state",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "dig_hole": {
                  "type": "boolean"
                },
                "drop_cache": {
                  "type": "boolean"
                },
                "load_cache": {
                  "type": "boolean"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/snapshots/{ssId}/mincore": {
      "get": {
        "description": "Get mincore state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Mincore state",
            "schema": {
              "type": "object",
              "properties": {
                "n_nz_regions": {
                  "type": "integer"
                },
                "n_ws_regions": {
                  "type": "integer"
                },
                "nlayers": {
                  "type": "integer"
                },
                "nz_region_size": {
                  "type": "integer"
                },
                "ws_region_size": {
                  "type": "integer"
                }
              }
            }
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      },
      "put": {
        "description": "Put mincore state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "source",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      },
      "post": {
        "description": "Add mincore layer",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "name": "layer",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "fromDiff": {
                  "type": "string"
                },
                "position": {
                  "type": "integer"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      },
      "patch": {
        "description": "Change mincore state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "name": "state",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "drop_ws_cache": {
                  "type": "boolean"
                },
                "from_records_size": {
                  "type": "integer"
                },
                "inactive_ws": {
                  "type": "boolean"
                },
                "interval_threshold": {
                  "type": "integer"
                },
                "mincore_cache": {
                  "type": "array",
                  "items": {
                    "type": "integer"
                  }
                },
                "size_threshold": {
                  "type": "integer"
                },
                "to_ws_file": {
                  "type": "string"
                },
                "trim_regions": {
                  "type": "boolean"
                },
                "zero_ws": {
                  "type": "boolean"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/snapshots/{ssId}/reap": {
      "get": {
        "description": "get reap state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      },
      "delete": {
        "description": "delete reap state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      },
      "patch": {
        "description": "Change reap state",
        "parameters": [
          {
            "type": "string",
            "name": "ssId",
            "in": "path",
            "required": true
          },
          {
            "name": "cache",
            "in": "body",
            "schema": {
              "type": "boolean"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/ui": {
      "get": {
        "description": "UI",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/ui/data": {
      "get": {
        "description": "UI",
        "produces": [
          "application/json"
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/vmms": {
      "post": {
        "description": "Create a VMM in the pool",
        "parameters": [
          {
            "name": "VMM",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "enableReap": {
                  "type": "boolean"
                },
                "namespace": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VM"
            }
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/vms": {
      "get": {
        "description": "Returns a list of active VMs",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/VM"
              }
            }
          }
        }
      },
      "post": {
        "description": "Create a new VM",
        "parameters": [
          {
            "name": "VM",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "func_name": {
                  "type": "string"
                },
                "namespace": {
                  "type": "string"
                },
                "ssId": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VM"
            }
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    },
    "/vms/{vmId}": {
      "get": {
        "description": "Describe a VM",
        "parameters": [
          {
            "type": "string",
            "name": "vmId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/VM"
            }
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      },
      "delete": {
        "description": "Stop a VM",
        "parameters": [
          {
            "type": "string",
            "name": "vmId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Invalid request",
            "schema": {
              "type": "object",
              "properties": {
                "message": {
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Function": {
      "type": "object",
      "required": [
        "func_name"
      ],
      "properties": {
        "func_name": {
          "type": "string"
        },
        "image": {
          "type": "string"
        },
        "kernel": {
          "type": "string"
        },
        "mem_size": {
          "type": "integer"
        },
        "vcpu": {
          "type": "integer"
        }
      }
    },
    "Invocation": {
      "type": "object",
      "required": [
        "func_name"
      ],
      "properties": {
        "enableReap": {
          "type": "boolean"
        },
        "func_name": {
          "type": "string"
        },
        "loadMincore": {
          "type": "array",
          "items": {
            "type": "integer"
          }
        },
        "mincore": {
          "type": "integer",
          "default": -1
        },
        "mincore_size": {
          "type": "integer"
        },
        "namespace": {
          "type": "string"
        },
        "overlay_regions": {
          "type": "boolean"
        },
        "params": {
          "type": "string"
        },
        "ssId": {
          "type": "string"
        },
        "use_mem_file": {
          "type": "boolean"
        },
        "use_ws_file": {
          "type": "boolean"
        },
        "vmId": {
          "type": "string"
        },
        "vmm_load_ws": {
          "type": "boolean"
        },
        "wsFileDirectIo": {
          "type": "boolean"
        },
        "wsSingleRead": {
          "type": "boolean"
        }
      }
    },
    "Snapshot": {
      "type": "object",
      "required": [
        "vmId"
      ],
      "properties": {
        "interval_threshold": {
          "type": "integer"
        },
        "mem_file_path": {
          "type": "string"
        },
        "record_regions": {
          "type": "boolean"
        },
        "size_threshold": {
          "type": "integer"
        },
        "snapshot_path": {
          "type": "string"
        },
        "snapshot_type": {
          "type": "string"
        },
        "ssId": {
          "type": "string"
        },
        "version": {
          "type": "string"
        },
        "vmId": {
          "type": "string"
        }
      }
    },
    "VM": {
      "type": "object",
      "required": [
        "vmId"
      ],
      "properties": {
        "state": {
          "type": "string"
        },
        "vmConf": {
          "type": "object"
        },
        "vmId": {
          "type": "string"
        },
        "vmPath": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "400Error": {
      "description": "Invalid request",
      "schema": {
        "type": "object",
        "properties": {
          "message": {
            "type": "string"
          }
        }
      }
    }
  }
}`))
}
