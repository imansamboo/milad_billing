<?php

namespace App\Controller\Admin;

use App\Entity\Ticket;
use App\Entity\TicketPost;
use App\Entity\User;
use App\Repository\UserRepository;
use EasyCorp\Bundle\EasyAdminBundle\Controller\AbstractCrudController;
use EasyCorp\Bundle\EasyAdminBundle\Field\DateTimeField;
use EasyCorp\Bundle\EasyAdminBundle\Field\Field;
use EasyCorp\Bundle\EasyAdminBundle\Field\IntegerField;
use EasyCorp\Bundle\EasyAdminBundle\Field\MoneyField;
use EasyCorp\Bundle\EasyAdminBundle\Field\TextEditorField;
use EasyCorp\Bundle\EasyAdminBundle\Field\TextField;
use EasyCorp\Bundle\EasyAdminBundle\Field\IdField;
use EasyCorp\Bundle\EasyAdminBundle\Field\ChoiceField;
use EasyCorp\Bundle\EasyAdminBundle\Field\HiddenField;

class TicketPostCrudController extends AbstractCrudController
{
    public static function getEntityFqcn(): string
    {
        return TicketPost::class;
    }


    public function configureFields(string $pageName): iterable
    {
        /* user, ticket, content, created_at, updated_at*/
        $user_repo = $this->getDoctrine()->getRepository(User::class);
        $user_choices =  $user_repo->getUserChoices();
        $ticket_choices = $this->getDoctrine()->getRepository(Ticket::class)->getFavoriteChoices();
        $content = TextEditorField::new('content');
        $user = ChoiceField::new("user", "creator")->setChoices($user_choices)->onlyOnForms();
        $ticket = ChoiceField::new("ticket", "related_ticket")->setChoices($ticket_choices)->onlyOnForms();
        $username = Field::new("owner", "Owner")->onlyOnIndex();
        return [
            $content,
            $user,
            $ticket,
            $username
        ];
    }

}
