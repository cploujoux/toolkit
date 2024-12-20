from typing import Any, TypeVar, Union

from attrs import define as _attrs_define
from attrs import field as _attrs_field

from ..types import UNSET, Unset

T = TypeVar("T", bound="IncreaseAndRateMetric")


@_attrs_define
class IncreaseAndRateMetric:
    """Metrics for resources

    Attributes:
        inference_global (Union[Unset, Any]): Historical requests for all resources globally
        query (Union[Unset, Any]): Number of requests for all resources globally
    """

    inference_global: Union[Unset, Any] = UNSET
    query: Union[Unset, Any] = UNSET
    additional_properties: dict[str, Any] = _attrs_field(init=False, factory=dict)

    def to_dict(self) -> dict[str, Any]:
        inference_global = self.inference_global

        query = self.query

        field_dict: dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if inference_global is not UNSET:
            field_dict["inferenceGlobal"] = inference_global
        if query is not UNSET:
            field_dict["query"] = query

        return field_dict

    @classmethod
    def from_dict(cls: type[T], src_dict: dict[str, Any]) -> T:
        if not src_dict:
            return None
        d = src_dict.copy()
        inference_global = d.pop("inferenceGlobal", UNSET)

        query = d.pop("query", UNSET)

        increase_and_rate_metric = cls(
            inference_global=inference_global,
            query=query,
        )

        increase_and_rate_metric.additional_properties = d
        return increase_and_rate_metric

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
