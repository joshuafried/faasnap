# coding: utf-8

"""
    faasnap

    FaaSnap API  # noqa: E501

    OpenAPI spec version: 1.0.0
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from swagger_client.configuration import Configuration


class VM(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'vm_id': 'str',
        'state': 'str',
        'vm_conf': 'object',
        'vm_path': 'str'
    }

    attribute_map = {
        'vm_id': 'vmId',
        'state': 'state',
        'vm_conf': 'vmConf',
        'vm_path': 'vmPath'
    }

    def __init__(self, vm_id=None, state=None, vm_conf=None, vm_path=None, _configuration=None):  # noqa: E501
        """VM - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._vm_id = None
        self._state = None
        self._vm_conf = None
        self._vm_path = None
        self.discriminator = None

        self.vm_id = vm_id
        if state is not None:
            self.state = state
        if vm_conf is not None:
            self.vm_conf = vm_conf
        if vm_path is not None:
            self.vm_path = vm_path

    @property
    def vm_id(self):
        """Gets the vm_id of this VM.  # noqa: E501


        :return: The vm_id of this VM.  # noqa: E501
        :rtype: str
        """
        return self._vm_id

    @vm_id.setter
    def vm_id(self, vm_id):
        """Sets the vm_id of this VM.


        :param vm_id: The vm_id of this VM.  # noqa: E501
        :type: str
        """
        if self._configuration.client_side_validation and vm_id is None:
            raise ValueError("Invalid value for `vm_id`, must not be `None`")  # noqa: E501

        self._vm_id = vm_id

    @property
    def state(self):
        """Gets the state of this VM.  # noqa: E501


        :return: The state of this VM.  # noqa: E501
        :rtype: str
        """
        return self._state

    @state.setter
    def state(self, state):
        """Sets the state of this VM.


        :param state: The state of this VM.  # noqa: E501
        :type: str
        """

        self._state = state

    @property
    def vm_conf(self):
        """Gets the vm_conf of this VM.  # noqa: E501


        :return: The vm_conf of this VM.  # noqa: E501
        :rtype: object
        """
        return self._vm_conf

    @vm_conf.setter
    def vm_conf(self, vm_conf):
        """Sets the vm_conf of this VM.


        :param vm_conf: The vm_conf of this VM.  # noqa: E501
        :type: object
        """

        self._vm_conf = vm_conf

    @property
    def vm_path(self):
        """Gets the vm_path of this VM.  # noqa: E501


        :return: The vm_path of this VM.  # noqa: E501
        :rtype: str
        """
        return self._vm_path

    @vm_path.setter
    def vm_path(self, vm_path):
        """Sets the vm_path of this VM.


        :param vm_path: The vm_path of this VM.  # noqa: E501
        :type: str
        """

        self._vm_path = vm_path

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(VM, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, VM):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, VM):
            return True

        return self.to_dict() != other.to_dict()
