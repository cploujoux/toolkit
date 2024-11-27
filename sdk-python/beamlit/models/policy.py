from typing import TYPE_CHECKING, Any, Dict, List, TypeVar, Union, cast

from attrs import define as _attrs_define
from attrs import field as _attrs_field

from ..types import UNSET, Unset

if TYPE_CHECKING:
    from ..models.flavor import Flavor
    from ..models.labels import Labels
    from ..models.policy_location import PolicyLocation


T = TypeVar("T", bound="Policy")


@_attrs_define
class Policy:
    """Rule that controls how a deployment is made and served (e.g. location restrictions)

    Attributes:
        created_at (Union[Unset, str]): The date and time when the resource was created
        created_by (Union[Unset, str]): The user or service account who created the resource
        updated_at (Union[Unset, str]): The date and time when the resource was updated
        updated_by (Union[Unset, str]): The user or service account who updated the resource
        display_name (Union[Unset, str]): Policy display name
        flavors (Union[Unset, List['Flavor']]): Types of hardware available for deployments
        labels (Union[Unset, Labels]): Labels
        locations (Union[Unset, List['PolicyLocation']]): PolicyLocations is a local type that wraps a slice of Location
        name (Union[Unset, str]): Policy name
        resource_types (Union[Unset, List[str]]): PolicyResourceTypes is a local type that wraps a slice of
            PolicyResourceType
        type (Union[Unset, str]): Policy type, can be location or flavor
        workspace (Union[Unset, str]): The workspace the policy belongs to
    """

    created_at: Union[Unset, str] = UNSET
    created_by: Union[Unset, str] = UNSET
    updated_at: Union[Unset, str] = UNSET
    updated_by: Union[Unset, str] = UNSET
    display_name: Union[Unset, str] = UNSET
    flavors: Union[Unset, List["Flavor"]] = UNSET
    labels: Union[Unset, "Labels"] = UNSET
    locations: Union[Unset, List["PolicyLocation"]] = UNSET
    name: Union[Unset, str] = UNSET
    resource_types: Union[Unset, List[str]] = UNSET
    type: Union[Unset, str] = UNSET
    workspace: Union[Unset, str] = UNSET
    additional_properties: dict[str, Any] = _attrs_field(init=False, factory=dict)

    def to_dict(self) -> dict[str, Any]:
        created_at = self.created_at

        created_by = self.created_by

        updated_at = self.updated_at

        updated_by = self.updated_by

        display_name = self.display_name

        flavors: Union[Unset, List[Dict[str, Any]]] = UNSET
        if not isinstance(self.flavors, Unset):
            flavors = []
            for componentsschemas_flavors_item_data in self.flavors:
                componentsschemas_flavors_item = componentsschemas_flavors_item_data.to_dict()
                flavors.append(componentsschemas_flavors_item)

        labels: Union[Unset, Dict[str, Any]] = UNSET
        if not isinstance(self.labels, Unset):
            labels = self.labels.to_dict()

        locations: Union[Unset, List[Dict[str, Any]]] = UNSET
        if not isinstance(self.locations, Unset):
            locations = []
            for componentsschemas_policy_locations_item_data in self.locations:
                componentsschemas_policy_locations_item = componentsschemas_policy_locations_item_data.to_dict()
                locations.append(componentsschemas_policy_locations_item)

        name = self.name

        resource_types: Union[Unset, List[str]] = UNSET
        if not isinstance(self.resource_types, Unset):
            resource_types = self.resource_types

        type = self.type

        workspace = self.workspace

        field_dict: dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if created_at is not UNSET:
            field_dict["created_at"] = created_at
        if created_by is not UNSET:
            field_dict["created_by"] = created_by
        if updated_at is not UNSET:
            field_dict["updated_at"] = updated_at
        if updated_by is not UNSET:
            field_dict["updated_by"] = updated_by
        if display_name is not UNSET:
            field_dict["display_name"] = display_name
        if flavors is not UNSET:
            field_dict["flavors"] = flavors
        if labels is not UNSET:
            field_dict["labels"] = labels
        if locations is not UNSET:
            field_dict["locations"] = locations
        if name is not UNSET:
            field_dict["name"] = name
        if resource_types is not UNSET:
            field_dict["resource_types"] = resource_types
        if type is not UNSET:
            field_dict["type"] = type
        if workspace is not UNSET:
            field_dict["workspace"] = workspace

        return field_dict

    @classmethod
    def from_dict(cls: type[T], src_dict: dict[str, Any]) -> T:
        from ..models.flavor import Flavor
        from ..models.labels import Labels
        from ..models.policy_location import PolicyLocation

        d = src_dict.copy()
        created_at = d.pop("created_at", UNSET)

        created_by = d.pop("created_by", UNSET)

        updated_at = d.pop("updated_at", UNSET)

        updated_by = d.pop("updated_by", UNSET)

        display_name = d.pop("display_name", UNSET)

        flavors = []
        _flavors = d.pop("flavors", UNSET)
        for componentsschemas_flavors_item_data in _flavors or []:
            componentsschemas_flavors_item = Flavor.from_dict(componentsschemas_flavors_item_data)

            flavors.append(componentsschemas_flavors_item)

        _labels = d.pop("labels", UNSET)
        labels: Union[Unset, Labels]
        if isinstance(_labels, Unset):
            labels = UNSET
        else:
            labels = Labels.from_dict(_labels)

        locations = []
        _locations = d.pop("locations", UNSET)
        for componentsschemas_policy_locations_item_data in _locations or []:
            componentsschemas_policy_locations_item = PolicyLocation.from_dict(
                componentsschemas_policy_locations_item_data
            )

            locations.append(componentsschemas_policy_locations_item)

        name = d.pop("name", UNSET)

        resource_types = cast(List[str], d.pop("resource_types", UNSET))

        type = d.pop("type", UNSET)

        workspace = d.pop("workspace", UNSET)

        policy = cls(
            created_at=created_at,
            created_by=created_by,
            updated_at=updated_at,
            updated_by=updated_by,
            display_name=display_name,
            flavors=flavors,
            labels=labels,
            locations=locations,
            name=name,
            resource_types=resource_types,
            type=type,
            workspace=workspace,
        )

        policy.additional_properties = d
        return policy

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
