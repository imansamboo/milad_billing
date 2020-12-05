<?php

namespace App\Entity;

use App\Repository\CdrRepository;
use Doctrine\ORM\Mapping as ORM;

/**
 * @ORM\Entity(repositoryClass=CdrRepository::class)
 */
class Cdr
{
    /**
     * @ORM\Id
     * @ORM\GeneratedValue
     * @ORM\Column(type="integer")
     */
    private $id;

    /**
     * @ORM\Column(type="string", length=40, nullable=true)
     */
    private $sim_number;

    /**
     * @ORM\Column(type="string", length=40, nullable=true)
     */
    private $account_number;

    /**
     * @ORM\Column(type="string", length=40, nullable=true)
     */
    private $invoice_number;

    /**
     * @ORM\Column(type="string", length=40, nullable=true)
     */
    private $ocean_region;

    /**
     * @ORM\Column(type="date", nullable=true)
     */
    private $date;

    /**
     * @ORM\Column(type="time", nullable=true)
     */
    private $time;

    /**
     * @ORM\Column(type="string", length=40, nullable=true)
     */
    private $originator_number;

    /**
     * @ORM\Column(type="string", length=40, nullable=true)
     */
    private $subscriber;

    /**
     * @ORM\Column(type="string", length=40)
     */
    private $destination_number;

    /**
     * @ORM\Column(type="float")
     */
    private $volume;

    /**
     * @ORM\Column(type="string", length=10)
     */
    private $unit;

    /**
     * @ORM\Column(type="string", length=40, nullable=true)
     */
    private $rate;

    /**
     * @ORM\Column(type="float")
     */
    private $total_charge;

    /**
     * @ORM\Column(type="string", length=80, nullable=true)
     */
    private $equipment_type;

    /**
     * @ORM\Column(type="string", length=80, nullable=true)
     */
    private $call_2_call_voice;

    /**
     * @ORM\Column(type="integer")
     */
    private $call_identifier_id;

    /**
     * @ORM\Column(type="string", length=40)
     */
    private $originator_country;

    /**
     * @ORM\Column(type="string", length=40)
     */
    private $destination_country;

    /**
     * @ORM\Column(type="string", length=40)
     */
    private $provider;

    /**
     * @ORM\Column(type="integer")
     */
    private $upstream_rate;

    /**
     * @ORM\Column(type="integer")
     */
    private $downstream_rate;

    /**
     * @ORM\Column(type="string", length=40)
     */
    private $data_session_id;

    /**
     * @ORM\Column(type="string", length=255)
     */
    private $apn;

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getSimNumber(): ?string
    {
        return $this->sim_number;
    }

    public function setSimNumber(?int $sim_number): self
    {
        $this->sim_number = $sim_number;

        return $this;
    }

    public function getAccountNumber(): ?string
    {
        return $this->account_number;
    }

    public function setAccountNumber(?string $account_number): self
    {
        $this->account_number = $account_number;

        return $this;
    }

    public function getInvoiceNumber(): ?string
    {
        return $this->invoice_number;
    }

    public function setInvoiceNumber(?string $invoice_number): self
    {
        $this->invoice_number = $invoice_number;

        return $this;
    }

    public function getOceanRegion(): ?string
    {
        return $this->ocean_region;
    }

    public function setOceanRegion(?string $ocean_region): self
    {
        $this->ocean_region = $ocean_region;

        return $this;
    }

    public function getDate(): ?\DateTimeInterface
    {
        return $this->date;
    }

    public function setDate(?\DateTimeInterface $date): self
    {
        $this->date = $date;

        return $this;
    }

    public function getTime(): ?\DateTimeInterface
    {
        return $this->time;
    }

    public function setTime(?\DateTimeInterface $time): self
    {
        $this->time = $time;

        return $this;
    }

    public function getOriginatorNumber(): ?string
    {
        return $this->originator_number;
    }

    public function setOriginatorNumber(?int $originator_number): self
    {
        $this->originator_number = $originator_number;

        return $this;
    }

    public function getSubscriber(): ?string
    {
        return $this->subscriber;
    }

    public function setSubscriber(?string $subscriber): self
    {
        $this->subscriber = $subscriber;

        return $this;
    }

    public function getDestinationNumber(): ?string
    {
        return $this->destination_number;
    }

    public function setDestinationNumber(string $destination_number): self
    {
        $this->destination_number = $destination_number;

        return $this;
    }

    public function getVolume(): ?float
    {
        return $this->volume;
    }

    public function setVolume(float $volume): self
    {
        $this->volume = $volume;

        return $this;
    }

    public function getUnit(): ?string
    {
        return $this->unit;
    }

    public function setUnit(string $unit): self
    {
        $this->unit = $unit;

        return $this;
    }

    public function getRate(): ?string
    {
        return $this->rate;
    }

    public function setRate(?string $rate): self
    {
        $this->rate = $rate;

        return $this;
    }

    public function getTotalCharge(): ?float
    {
        return $this->total_charge;
    }

    public function setTotalCharge(float $total_charge): self
    {
        $this->total_charge = $total_charge;

        return $this;
    }

    public function getEquipmentType(): ?string
    {
        return $this->equipment_type;
    }

    public function setEquipmentType(?string $equipment_type): self
    {
        $this->equipment_type = $equipment_type;

        return $this;
    }

    public function getCall2CallVoice(): ?string
    {
        return $this->call_2_call_voice;
    }

    public function setCall2CallVoice(?string $call_2_call_voice): self
    {
        $this->call_2_call_voice = $call_2_call_voice;

        return $this;
    }

    public function getCallIdentifierId(): ?int
    {
        return $this->call_identifier_id;
    }

    public function setCallIdentifierId(int $call_identifier_id): self
    {
        $this->call_identifier_id = $call_identifier_id;

        return $this;
    }

    public function getOriginatorCountry(): ?string
    {
        return $this->originator_country;
    }

    public function setOriginatorCountry(string $originator_country): self
    {
        $this->originator_country = $originator_country;

        return $this;
    }

    public function getDestinationCountry(): ?string
    {
        return $this->destination_country;
    }

    public function setDestinationCountry(string $destination_country): self
    {
        $this->destination_country = $destination_country;

        return $this;
    }

    public function getProvider(): ?string
    {
        return $this->provider;
    }

    public function setProvider(string $provider): self
    {
        $this->provider = $provider;

        return $this;
    }

    public function getUpstreamRate(): ?int
    {
        return $this->upstream_rate;
    }

    public function setUpstreamRate(int $upstream_rate): self
    {
        $this->upstream_rate = $upstream_rate;

        return $this;
    }

    public function getDownstreamRate(): ?int
    {
        return $this->downstream_rate;
    }

    public function setDownstreamRate(int $downstream_rate): self
    {
        $this->downstream_rate = $downstream_rate;

        return $this;
    }

    public function getDataSessionId(): ?string
    {
        return $this->data_session_id;
    }

    public function setDataSessionId(string $data_session_id): self
    {
        $this->data_session_id = $data_session_id;

        return $this;
    }

    public function getApn(): ?string
    {
        return $this->apn;
    }

    public function setApn(string $apn): self
    {
        $this->apn = $apn;

        return $this;
    }
}
