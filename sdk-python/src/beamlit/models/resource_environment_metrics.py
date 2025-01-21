from typing import TYPE_CHECKING, Any, TypeVar, Union

from attrs import define as _attrs_define
from attrs import field as _attrs_field

from ..types import UNSET, Unset

if TYPE_CHECKING:
    from ..models.latency_metric import LatencyMetric
    from ..models.metric import Metric


T = TypeVar("T", bound="ResourceEnvironmentMetrics")


@_attrs_define
class ResourceEnvironmentMetrics:
    """Metrics for a single resource deployment (eg. model deployment, function deployment)

    Attributes:
        inference_global (Union[Unset, list['Metric']]): Array of metrics
        last_n_requests (Union[Unset, list['Metric']]): Array of metrics
        latency (Union[Unset, LatencyMetric]): Latency metrics
        request_total (Union[Unset, Any]): Number of requests for the resource globally
        request_total_per_code (Union[Unset, Any]): Number of requests for the resource globally per code
        rps (Union[Unset, Any]): Number of requests per second for the resource globally
        rps_per_code (Union[Unset, Any]): Number of requests per second for the resource globally per code
    """

    inference_global: Union[Unset, list["Metric"]] = UNSET
    last_n_requests: Union[Unset, list["Metric"]] = UNSET
    latency: Union[Unset, "LatencyMetric"] = UNSET
    request_total: Union[Unset, Any] = UNSET
    request_total_per_code: Union[Unset, Any] = UNSET
    rps: Union[Unset, Any] = UNSET
    rps_per_code: Union[Unset, Any] = UNSET
    additional_properties: dict[str, Any] = _attrs_field(init=False, factory=dict)

    def to_dict(self) -> dict[str, Any]:
        inference_global: Union[Unset, list[dict[str, Any]]] = UNSET
        if not isinstance(self.inference_global, Unset):
            inference_global = []
            for componentsschemas_array_metric_item_data in self.inference_global:
                componentsschemas_array_metric_item = componentsschemas_array_metric_item_data.to_dict()
                inference_global.append(componentsschemas_array_metric_item)

        last_n_requests: Union[Unset, list[dict[str, Any]]] = UNSET
        if not isinstance(self.last_n_requests, Unset):
            last_n_requests = []
            for componentsschemas_array_metric_item_data in self.last_n_requests:
                componentsschemas_array_metric_item = componentsschemas_array_metric_item_data.to_dict()
                last_n_requests.append(componentsschemas_array_metric_item)

        latency: Union[Unset, dict[str, Any]] = UNSET
        if not isinstance(self.latency, Unset):
            latency = self.latency.to_dict()

        request_total = self.request_total

        request_total_per_code = self.request_total_per_code

        rps = self.rps

        rps_per_code = self.rps_per_code

        field_dict: dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if inference_global is not UNSET:
            field_dict["inferenceGlobal"] = inference_global
        if last_n_requests is not UNSET:
            field_dict["lastNRequests"] = last_n_requests
        if latency is not UNSET:
            field_dict["latency"] = latency
        if request_total is not UNSET:
            field_dict["requestTotal"] = request_total
        if request_total_per_code is not UNSET:
            field_dict["requestTotalPerCode"] = request_total_per_code
        if rps is not UNSET:
            field_dict["rps"] = rps
        if rps_per_code is not UNSET:
            field_dict["rpsPerCode"] = rps_per_code

        return field_dict

    @classmethod
    def from_dict(cls: type[T], src_dict: dict[str, Any]) -> T:
        from ..models.latency_metric import LatencyMetric
        from ..models.metric import Metric

        if not src_dict:
            return None
        d = src_dict.copy()
        inference_global = []
        _inference_global = d.pop("inferenceGlobal", UNSET)
        for componentsschemas_array_metric_item_data in _inference_global or []:
            componentsschemas_array_metric_item = Metric.from_dict(componentsschemas_array_metric_item_data)

            inference_global.append(componentsschemas_array_metric_item)

        last_n_requests = []
        _last_n_requests = d.pop("lastNRequests", UNSET)
        for componentsschemas_array_metric_item_data in _last_n_requests or []:
            componentsschemas_array_metric_item = Metric.from_dict(componentsschemas_array_metric_item_data)

            last_n_requests.append(componentsschemas_array_metric_item)

        _latency = d.pop("latency", UNSET)
        latency: Union[Unset, LatencyMetric]
        if isinstance(_latency, Unset):
            latency = UNSET
        else:
            latency = LatencyMetric.from_dict(_latency)

        request_total = d.pop("requestTotal", UNSET)

        request_total_per_code = d.pop("requestTotalPerCode", UNSET)

        rps = d.pop("rps", UNSET)

        rps_per_code = d.pop("rpsPerCode", UNSET)

        resource_environment_metrics = cls(
            inference_global=inference_global,
            last_n_requests=last_n_requests,
            latency=latency,
            request_total=request_total,
            request_total_per_code=request_total_per_code,
            rps=rps,
            rps_per_code=rps_per_code,
        )

        resource_environment_metrics.additional_properties = d
        return resource_environment_metrics

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
