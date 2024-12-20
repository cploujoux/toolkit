from typing import Any, TypeVar, Union

from attrs import define as _attrs_define
from attrs import field as _attrs_field

from ..types import UNSET, Unset

T = TypeVar("T", bound="CoreStatus")


@_attrs_define
class CoreStatus:
    """Core status

    Attributes:
        deployment_status (Union[Unset, str]): The status of the core, can be CREATED, UPDATED, DELETED, DEPLOYED,
            DISABLED, or FAILED
    """

    deployment_status: Union[Unset, str] = UNSET
    additional_properties: dict[str, Any] = _attrs_field(init=False, factory=dict)

    def to_dict(self) -> dict[str, Any]:
        deployment_status = self.deployment_status

        field_dict: dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if deployment_status is not UNSET:
            field_dict["deploymentStatus"] = deployment_status

        return field_dict

    @classmethod
    def from_dict(cls: type[T], src_dict: dict[str, Any]) -> T:
        if not src_dict:
            return None
        d = src_dict.copy()
        deployment_status = d.pop("deploymentStatus", UNSET)

        core_status = cls(
            deployment_status=deployment_status,
        )

        core_status.additional_properties = d
        return core_status

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
