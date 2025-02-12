from typing import Any, TypeVar, Union, cast

from attrs import define as _attrs_define
from attrs import field as _attrs_field

from ..types import UNSET, Unset

T = TypeVar("T", bound="AgentInformationRequest")


@_attrs_define
class AgentInformationRequest:
    """generation agent information request

    Attributes:
        functions (Union[Unset, list[Any]]): Functions to generate information for
    """

    functions: Union[Unset, list[Any]] = UNSET
    additional_properties: dict[str, Any] = _attrs_field(init=False, factory=dict)

    def to_dict(self) -> dict[str, Any]:
        functions: Union[Unset, list[Any]] = UNSET
        if not isinstance(self.functions, Unset):
            functions = self.functions

        field_dict: dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if functions is not UNSET:
            field_dict["functions"] = functions

        return field_dict

    @classmethod
    def from_dict(cls: type[T], src_dict: dict[str, Any]) -> T:
        if not src_dict:
            return None
        d = src_dict.copy()
        functions = cast(list[Any], d.pop("functions", UNSET))

        agent_information_request = cls(
            functions=functions,
        )

        agent_information_request.additional_properties = d
        return agent_information_request

    @property
    def additional_keys(self) -> list[str]:
        return list(self.additional_properties.keys())

    def __getitem__(self, key: str) -> Any:
        return self.additional_properties[key]

    def __setitem__(self, key: str, value: Any) -> None:
        self.additional_properties[key] = value

    def __delitem__(self, key: str) -> None:
        del self.additional_properties[key]

    def __contains__(self, key: str) -> bool:
        return key in self.additional_properties
