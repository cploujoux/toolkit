from typing import Any, TypeVar, Union

from attrs import define as _attrs_define
from attrs import field as _attrs_field

from ..types import UNSET, Unset

T = TypeVar("T", bound="BillingAddress")


@_attrs_define
class BillingAddress:
    """BillingAddress

    Attributes:
        city (Union[Unset, str]): City
        country (Union[Unset, str]): Country
        first_name (Union[Unset, str]): First name
        last_name (Union[Unset, str]): Last name
        phone_number (Union[Unset, str]): Phone number
        postal_code (Union[Unset, str]): Postal code
        state (Union[Unset, str]): State/Province/Region
        street1 (Union[Unset, str]): Street 1
        street2 (Union[Unset, str]): Street 2
    """

    city: Union[Unset, str] = UNSET
    country: Union[Unset, str] = UNSET
    first_name: Union[Unset, str] = UNSET
    last_name: Union[Unset, str] = UNSET
    phone_number: Union[Unset, str] = UNSET
    postal_code: Union[Unset, str] = UNSET
    state: Union[Unset, str] = UNSET
    street1: Union[Unset, str] = UNSET
    street2: Union[Unset, str] = UNSET
    additional_properties: dict[str, Any] = _attrs_field(init=False, factory=dict)

    def to_dict(self) -> dict[str, Any]:
        city = self.city

        country = self.country

        first_name = self.first_name

        last_name = self.last_name

        phone_number = self.phone_number

        postal_code = self.postal_code

        state = self.state

        street1 = self.street1

        street2 = self.street2

        field_dict: dict[str, Any] = {}
        field_dict.update(self.additional_properties)
        field_dict.update({})
        if city is not UNSET:
            field_dict["city"] = city
        if country is not UNSET:
            field_dict["country"] = country
        if first_name is not UNSET:
            field_dict["firstName"] = first_name
        if last_name is not UNSET:
            field_dict["lastName"] = last_name
        if phone_number is not UNSET:
            field_dict["phoneNumber"] = phone_number
        if postal_code is not UNSET:
            field_dict["postalCode"] = postal_code
        if state is not UNSET:
            field_dict["state"] = state
        if street1 is not UNSET:
            field_dict["street1"] = street1
        if street2 is not UNSET:
            field_dict["street2"] = street2

        return field_dict

    @classmethod
    def from_dict(cls: type[T], src_dict: dict[str, Any]) -> T:
        if not src_dict:
            return None
        d = src_dict.copy()
        city = d.pop("city", UNSET)

        country = d.pop("country", UNSET)

        first_name = d.pop("firstName", UNSET)

        last_name = d.pop("lastName", UNSET)

        phone_number = d.pop("phoneNumber", UNSET)

        postal_code = d.pop("postalCode", UNSET)

        state = d.pop("state", UNSET)

        street1 = d.pop("street1", UNSET)

        street2 = d.pop("street2", UNSET)

        billing_address = cls(
            city=city,
            country=country,
            first_name=first_name,
            last_name=last_name,
            phone_number=phone_number,
            postal_code=postal_code,
            state=state,
            street1=street1,
            street2=street2,
        )

        billing_address.additional_properties = d
        return billing_address

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
