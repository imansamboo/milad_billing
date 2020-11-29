<?php

namespace App\Entity;

use App\Repository\TicketRepository;
use Doctrine\Common\Collections\ArrayCollection;
use Doctrine\Common\Collections\Collection;
use Doctrine\ORM\Mapping as ORM;

/**
 * @ORM\Entity(repositoryClass=TicketRepository::class)
 */
class Ticket
{
    /**
     * @ORM\Id
     * @ORM\GeneratedValue
     * @ORM\Column(type="integer")
     */
    private $id;

    /**
     * @ORM\Column(type="string", length=40)
     */
    private $status;

    /**
     * @ORM\Column(type="integer")
     */
    private $owner;

    /**
     * @ORM\Column(type="datetime")
     */
    private $created_at;

    /**
     * @ORM\Column(type="datetime", nullable=true)
     */
    private $modified_at;

    /**
     * @ORM\Column(type="string", length=255)
     */
    private $subject;

    /**
     * @ORM\OneToMany(targetEntity=TicketPost::class, mappedBy="ticket", orphanRemoval=true)
     */
    private $ticketPosts;

    public function __construct()
    {
        $this->ticketPosts = new ArrayCollection();
    }

    public function getId(): ?int
    {
        return $this->id;
    }

    public function getStatus(): ?string
    {
        return $this->status;
    }

    public function setStatus(string $status): self
    {
        $this->status = $status;

        return $this;
    }

    public function getOwner(): ?int
    {
        return $this->owner;
    }

    public function setOwner(int $owner): self
    {
        $this->owner = $owner;

        return $this;
    }

    public function getCreatedAt(): ?\DateTimeInterface
    {
        return $this->created_at;
    }

    public function setCreatedAt(\DateTimeInterface $created_at): self
    {
        $this->created_at = $created_at;

        return $this;
    }

    public function getModifiedAt(): ?\DateTimeInterface
    {
        return $this->modified_at;
    }

    public function setModifiedAt(?\DateTimeInterface $modified_at): self
    {
        $this->modified_at = $modified_at;

        return $this;
    }

    public function getSubject(): ?string
    {
        return $this->subject;
    }

    public function setSubject(string $subject): self
    {
        $this->subject = $subject;

        return $this;
    }

    /**
     * @return Collection|TicketPost[]
     */
    public function getTicketPosts(): Collection
    {
        return $this->ticketPosts;
    }

    public function addTicketPost(TicketPost $ticketPost): self
    {
        if (!$this->ticketPosts->contains($ticketPost)) {
            $this->ticketPosts[] = $ticketPost;
            $ticketPost->setTicket($this);
        }

        return $this;
    }

    public function removeTicketPost(TicketPost $ticketPost): self
    {
        if ($this->ticketPosts->removeElement($ticketPost)) {
            // set the owning side to null (unless already changed)
            if ($ticketPost->getTicket() === $this) {
                $ticketPost->setTicket(null);
            }
        }

        return $this;
    }
}
