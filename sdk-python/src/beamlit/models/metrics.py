from typing import TYPE_CHECKING, Any, TypeVar, Union

from attrs import define as _attrs_define
from attrs import field as _attrs_field

from ..types import UNSET, Unset

if TYPE_CHECKING:
    from ..models.increase_and_rate_metric import IncreaseAndRateMetric


T = TypeVar("T", bound="Metrics")


@_attrs_define
class Metrics:
    """Metrics for resources

    Attributes:
        inference_global (Union[Unset, Any]): Historical requests for all resources globally
        query (Union[Unset, Any]): Number of requests for all resources globally
        agents (Union[Unset, IncreaseAndRateMetric]): Metrics for resources
        functions (Union[Unset, IncreaseAndRateMetric]): Metrics for resources
        models (Union[Unset, IncreaseAndRateMetric]): Metrics for resources
    """

    inference_global: Union[Unset, Any] = UNSET
    query: Union[Unset, Any] = UNSET
    agents: Union[Unset, "IncreaseAndRateMetric"] = UNSET
    functions: Union[Unset, "IncreaseAndRateMetric"] = UNSET
    models: Union[Unset, "IncreaseAndRateMetric"] = UNSET
    additional_properties: dict[str, Any] = _attrs_field(init=False, factory=dict)

    def to_dict(self) -> dict[str, Any]:
        inference_global = self.inference_global

        query = self.query

        agents: Union[Unset, dict[str, Any]] = UNSET
        if not isinstance(self.agents, Unset):
            agents = self.agents.to_dict()

        functions: Union[Unset, dict[str, Any]] = UNSET
        if not isinstance(self.functions, Unset):
            functions = self.functions.to_dict()

        models: Union[Unset, dict[str, Any]] = UNSET
        if not isinstance(self.models, Unset):
            models = self.models.to_dict()

        field_dict: dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if inference_global is not UNSET:
            field_dict["inferenceGlobal"] = inference_global
        if query is not UNSET:
            field_dict["query"] = query
        if agents is not UNSET:
            field_dict["agents"] = agents
        if functions is not UNSET:
            field_dict["functions"] = functions
        if models is not UNSET:
            field_dict["models"] = models

        return field_dict

    @classmethod
    def from_dict(cls: type[T], src_dict: dict[str, Any]) -> T:
        from ..models.increase_and_rate_metric import IncreaseAndRateMetric

        if not src_dict:
            return None
        d = src_dict.copy()
        inference_global = d.pop("inferenceGlobal", UNSET)

        query = d.pop("query", UNSET)

        _agents = d.pop("agents", UNSET)
        agents: Union[Unset, IncreaseAndRateMetric]
        if isinstance(_agents, Unset):
            agents = UNSET
        else:
            agents = IncreaseAndRateMetric.from_dict(_agents)

        _functions = d.pop("functions", UNSET)
        functions: Union[Unset, IncreaseAndRateMetric]
        if isinstance(_functions, Unset):
            functions = UNSET
        else:
            functions = IncreaseAndRateMetric.from_dict(_functions)

        _models = d.pop("models", UNSET)
        models: Union[Unset, IncreaseAndRateMetric]
        if isinstance(_models, Unset):
            models = UNSET
        else:
            models = IncreaseAndRateMetric.from_dict(_models)

        metrics = cls(
            inference_global=inference_global,
            query=query,
            agents=agents,
            functions=functions,
            models=models,
        )

        metrics.additional_properties = d
        return metrics

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
