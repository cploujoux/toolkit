from typing import TYPE_CHECKING, Any, Dict, List, TypeVar, Union

from attrs import define as _attrs_define
from attrs import field as _attrs_field

from ..types import UNSET, Unset

if TYPE_CHECKING:
    from ..models.labels import Labels
    from ..models.model_deployment import ModelDeployment


T = TypeVar("T", bound="ModelWithDeployments")


@_attrs_define
class ModelWithDeployments:
    """Logical object representing a model but with deployment definition inside

    Attributes:
        created_at (Union[Unset, str]): The date and time when the resource was created
        created_by (Union[Unset, str]): The user or service account who created the resource
        updated_at (Union[Unset, str]): The date and time when the resource was updated
        updated_by (Union[Unset, str]): The user or service account who updated the resource
        display_name (Union[Unset, str]): Model display name
        labels (Union[Unset, Labels]): Labels
        name (Union[Unset, str]): Model name
        workspace (Union[Unset, str]): The workspace the model belongs to
        deployments (Union[Unset, List['ModelDeployment']]):
    """

    created_at: Union[Unset, str] = UNSET
    created_by: Union[Unset, str] = UNSET
    updated_at: Union[Unset, str] = UNSET
    updated_by: Union[Unset, str] = UNSET
    display_name: Union[Unset, str] = UNSET
    labels: Union[Unset, "Labels"] = UNSET
    name: Union[Unset, str] = UNSET
    workspace: Union[Unset, str] = UNSET
    deployments: Union[Unset, List["ModelDeployment"]] = UNSET
    additional_properties: dict[str, Any] = _attrs_field(init=False, factory=dict)

    def to_dict(self) -> dict[str, Any]:
        created_at = self.created_at

        created_by = self.created_by

        updated_at = self.updated_at

        updated_by = self.updated_by

        display_name = self.display_name

        labels: Union[Unset, Dict[str, Any]] = UNSET
        if not isinstance(self.labels, Unset):
            labels = self.labels.to_dict()

        name = self.name

        workspace = self.workspace

        deployments: Union[Unset, List[Dict[str, Any]]] = UNSET
        if not isinstance(self.deployments, Unset):
            deployments = []
            for componentsschemas_model_deployments_item_data in self.deployments:
                componentsschemas_model_deployments_item = componentsschemas_model_deployments_item_data.to_dict()
                deployments.append(componentsschemas_model_deployments_item)

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
        if labels is not UNSET:
            field_dict["labels"] = labels
        if name is not UNSET:
            field_dict["name"] = name
        if workspace is not UNSET:
            field_dict["workspace"] = workspace
        if deployments is not UNSET:
            field_dict["deployments"] = deployments

        return field_dict

    @classmethod
    def from_dict(cls: type[T], src_dict: dict[str, Any]) -> T:
        from ..models.labels import Labels
        from ..models.model_deployment import ModelDeployment

        d = src_dict.copy()
        created_at = d.pop("created_at", UNSET)

        created_by = d.pop("created_by", UNSET)

        updated_at = d.pop("updated_at", UNSET)

        updated_by = d.pop("updated_by", UNSET)

        display_name = d.pop("display_name", UNSET)

        _labels = d.pop("labels", UNSET)
        labels: Union[Unset, Labels]
        if isinstance(_labels, Unset):
            labels = UNSET
        else:
            labels = Labels.from_dict(_labels)

        name = d.pop("name", UNSET)

        workspace = d.pop("workspace", UNSET)

        deployments = []
        _deployments = d.pop("deployments", UNSET)
        for componentsschemas_model_deployments_item_data in _deployments or []:
            componentsschemas_model_deployments_item = ModelDeployment.from_dict(
                componentsschemas_model_deployments_item_data
            )

            deployments.append(componentsschemas_model_deployments_item)

        model_with_deployments = cls(
            created_at=created_at,
            created_by=created_by,
            updated_at=updated_at,
            updated_by=updated_by,
            display_name=display_name,
            labels=labels,
            name=name,
            workspace=workspace,
            deployments=deployments,
        )

        model_with_deployments.additional_properties = d
        return model_with_deployments

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
